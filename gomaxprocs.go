package gomaxprocs

import (
	"os"
	"runtime"
)

// Get returns the current value of GOMAXPROCS.
func Get() int { return runtime.GOMAXPROCS(-1) }

// Set sets the value of GOMAXPROCS to n, unless the GOMAXPROCS environment variable was already set.
func Set(n int) {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(n)
	}
}

// SetToNumCPU sets GOMAXPROCS to the number of logical CPUs, unless the GOMAXPROCS environment variable was
// already set.
func SetToNumCPU() {
	Set(runtime.NumCPU())
}
