// Package belt provides a suite of quick utilities for writing go programs.
package belt

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	// Verbose should be set to true for debug output.
	Verbose = false
)

// Check terminates the program if the error is non-nil.
// The error is optionally logged if belt.Verbose is set to true.
func Check(err error) {
	if err != nil {
		Debug(err)

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

// HasPrefix returns true if and only if strings.HasPrefix(s, px) returns true for
// at least one px in prefixes.
func HasPrefix(s string, prefixes []string) bool {
	for _, px := range prefixes {
		if strings.HasPrefix(s, px) {
			return true
		}
	}

	return false
}

// HasSuffix returns true if and only if strings.HasSuffix(s, sx) returns true for
// at least one sx in suffixes.
func HasSuffix(s string, suffixes []string) bool {
	for _, sx := range suffixes {
		if strings.HasSuffix(s, sx) {
			return true
		}
	}

	return false
}

// Debug logs its arguments if belt.Verbose is true
func Debug(a ...interface{}) {
	if !Verbose {
		return
	}
	log.Println(a...)
}

// Debugf is a version of Debug that formats the output using fmt.Sprintf
func Debugf(format string, a ...interface{}) {
	Debug(fmt.Sprintf(format, a...))
}
