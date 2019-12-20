package main

import (
	"testing"

	"github.com/Workiva/go-datastructures/set"
	"github.com/stretchr/testify/assert"
)

func TestConfigValidate(t *testing.T) {
	cases := []struct {
		name   string
		input  *Config
		output *Config
		errs   []error
	}{
		{
			name:  "Should return errors with an empty input",
			input: &Config{},
			errs:  []error{errNoInputs, errNoTerms},
		},
		{
			name: "Should return errNoInputs if no inputs are given",
			input: &Config{
				Terms: set.New("foo"),
			},
			errs: []error{errNoInputs},
		},
		{
			name: "Should return errNoTerms if no terms are given",
			input: &Config{
				Inputs: []string{"./"},
			},
			errs: []error{errNoTerms},
		},
		{
			name: "Should not return errors if Inputs and Terms are given",
			input: &Config{
				Inputs: []string{"./"},
				Terms:  set.New("foo"),
			},
			errs: []error{},
		},
		{
			name: "Should not return errors if a full valid config is given",
			input: &Config{
				Inputs: []string{"./"},
				Terms:  set.New("foo"),
				Output: "./folder",
			},
			errs: []error{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			errs := tc.input.Validate()
			assert.ElementsMatch(t, tc.errs, errs, "Expected Config.Validate() to return all the correct errors")
		})
	}

	t.Run("Should set Output to `./output` if it is empty", func(t *testing.T) {
		input := &Config{
			Inputs: []string{"./"},
			Terms:  set.New("foo"),
		}
		expected := "./output"

		if assert.Empty(t, input.Validate()) {
			assert.Equal(t, expected, input.Output, "Expected Config.Validate() to set Output to `%s`", expected)
		}
	})
}
