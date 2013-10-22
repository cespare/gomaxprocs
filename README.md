# gomaxprocs

Gomaxprocs is a Go micro library that offers a slightly nicer interface to `runtime.GOMAXPROCS`.

* `Get() int`: Returns the current setting for `GOMAXPROCS`
* `Set(n int)`: Set a new value of `GOMAXPROCS`, but only if the `GOMAXPROCS` environment variable is not
  already set.
