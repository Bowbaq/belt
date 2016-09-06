// Package belt provides a suite of quick utilities for writing go programs.
package belt

import (
	"log"
	"os"
	"reflect"
)

var (
	// Verbose should be set to true for debug output.
	Verbose = false
)

// Check terminates the program if the error is non-nil.
// The error is optionally logged if belt.Verbose is set to true.
func Check(err error) {
	if err != nil {
		if Verbose {
			log.Println(err)
		}

		os.Exit(1)
	}
}

// Contains returns true if x is an element of xs. Though it is not enforced,
// x and the elements of xs must be of the same type for proper comparison.
// This uses a slow linear search suitable for small amounts of data. Equality
// between elements is tested using reflect.DeepEqual.
func Contains(xs interface{}, x interface{}) bool {
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if reflect.DeepEqual(x, slice.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}
