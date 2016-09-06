package belt

import (
	"log"
	"os"
)

var (
	// Set Verbose to true for debug output
	Verbose = false
)

// Check terminates the program if the error is non-nil.
// The error is optionally logged if belt.Verbose is set to true
func Check(err error) {
	if err != nil {
		if Verbose {
			log.Println(err)
		}

		os.Exit(1)
	}
}
