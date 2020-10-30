# go-tracked-error
Keep track of your errors while passing them through multiple functions. Very pragmatic library which enhances
the default error interface with a slice of tracks. Free of transitive dependencies.

## Usage
````go
package main

import (
	"errors"
	"github.com/ogermann/go-tracked-error"
	"log"
)

func main() {
	var err error

	err = terr.Track(errors.New("any error"))
	err = terr.Track(err, "with some additional", "text")

	log.Println(err)
}
````

This code will produce the following log message:
```
main.go:12: main_test.go:108: any error
    main.go:13: with some additional text
```