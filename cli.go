package main

import (
	"fmt"
	"os"

	"github.com/Workiva/go-datastructures/set"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/urfave/cli/v2"
)

// RunCLI is the entrypoint for the cli app
func RunCLI() error {
	app := &cli.App{
		Name:      "proto-filter",
		Usage:     "Filter out objects in a proto file based on a filter option",
		ArgsUsage: "[FILES]",
		Action:    action,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "include",
				Aliases: []string{"i"},
				Usage:   "`PATH` to add to the lookup path for proto files",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "`DIRECTORY` to emit the processed proto files to",
				Value:   "./output/",
			},
			&cli.StringSliceFlag{
				Name:     "term",
				Aliases:  []string{"t"},
				Usage:    "A `TERM` to filter for",
				Required: true,
			},
		},
	}

	return app.Run(os.Args)
}

func action(c *cli.Context) error {
	config := Config{
		Inputs:   c.Args().Slice(),
		Output:   c.String("output"),
		Includes: c.StringSlice("include"),
		Terms:    makeStringSet(c.StringSlice("term")),
	}

	if errs := config.Validate(); len(errs) != 0 {
		return fmt.Errorf("Invalid input: %s", errs)
	}

	parser := protoparse.Parser{
		ImportPaths:           config.Includes,
		InferImportPaths:      true,
		IncludeSourceCodeInfo: true,
	}
	descs, err := parser.ParseFiles(config.Inputs...)
	if err != nil {
		return err
	}

	output := make([]*desc.FileDescriptor, 0, len(descs))
	for _, fdesc := range descs {
		fileBuilder, err := builder.FromFile(fdesc)
		if err != nil {
			return err
		}
		if result, err := filterFile(fileBuilder, config.Terms); err != nil {
			return err
		} else if result {
			fDesc, err := fileBuilder.Build()
			if err != nil {
				return err
			}
			output = append(output, fDesc)
		}
	}

	printer := protoprint.Printer{}

	return printer.PrintProtosToFileSystem(output, config.Output)
}

// makeStringSet is a convenience wrapper which produces a new Set from a slice of strings
func makeStringSet(items []string) *set.Set {
	ifaceSlice := make([]interface{}, len(items))
	for i := range items {
		ifaceSlice[i] = items[i]
	}
	return set.New(ifaceSlice...)
}
