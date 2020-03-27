# System File Picker

This Go library provides access to platform-native file pickers from Go without requiring the use of C. The current implementation is very much a work in progress and should not be considered robust. The API is not yet stable, there is limited testing, and only Windows and Linux (GTK) are supported. In the future I hope to support Android, iOS, MacOS, the BSDs, and Solaris.

The API is dead simple, although users should be aware that this library is blocking - your program will halt while waiting for the selection to be made. If you'd like to make this an asynchronous process you can wrap the implementation with a goroutine as you would normally. Below is an example of usage:

```go
package main

import (
	"fmt"
	"log"
	"github.com/dylrich/sfp"
)

func main() {
	path, err := sfp.Pick()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)

	paths, err := sfp.PickMultiple()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range paths {
		fmt.Println(p)
	}
}
```
