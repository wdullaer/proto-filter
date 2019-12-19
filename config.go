package main

import (
	"errors"

	"github.com/Workiva/go-datastructures/set"
)

// Config encapsulates all configuration for the program
type Config struct {
	Inputs   []string
	Output   string
	Includes []string
	Terms    *set.Set
}

// Validate performs a limited set of validations on the configuration to make
// sure it represents an acceptable configuration
func (c *Config) Validate() []error {
	var errs = make([]error, 0, 5)

	if len(c.Inputs) == 0 {
		errs = append(errs, errors.New("No files given to process"))
	}

	if c.Terms.Len() == 0 {
		errs = append(errs, errors.New("No terms to filter for given"))
	}

	if len(c.Output) == 0 {
		c.Output = "./output"
	}

	return errs
}
