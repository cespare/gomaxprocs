package gomaxprocs

import (
	"os"
	"runtime"
)

func Get() int { return runtime.GOMAXPROCS(-1) }

func Set(n int) {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(n)
	}
}
