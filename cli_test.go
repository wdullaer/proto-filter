package main

import (
	"testing"

	"github.com/Workiva/go-datastructures/set"
	"github.com/stretchr/testify/assert"
)

func TestMakeStringSet(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		output *set.Set
	}{
		{
			name:   "Should convert an empty string slice",
			input:  []string{},
			output: set.New(),
		},
		{
			name:   "Should convert a string slice with no duplicates",
			input:  []string{"foo", "bar", "baz"},
			output: set.New("foo", "bar", "baz"),
		},
		{
			name:   "Should convert a string slice and remove the duplicates",
			input:  []string{"foo", "bar", "baz", "foo"},
			output: set.New("foo", "bar", "baz"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := makeStringSet(tc.input)
			assert.Equal(t, tc.output, result, "Expected `makeStringSet` with input %s to return %s", tc.input, tc.output)
		})
	}
}
