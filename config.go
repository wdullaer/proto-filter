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

var (
	errNoInputs = errors.New("No files given to process")
	errNoTerms  = errors.New("No terms given to filter for")
)

// Validate performs a limited set of validations on the configuration to make
// sure it represents an acceptable configuration
func (c *Config) Validate() []error {
	var errs = make([]error, 0, 5)

	if len(c.Inputs) == 0 {
		errs = append(errs, errNoInputs)
	}

	if c.Terms == nil || c.Terms.Len() == 0 {
		errs = append(errs, errNoTerms)
	}

	if len(c.Output) == 0 {
		c.Output = "./output"
	}

	return errs
}
