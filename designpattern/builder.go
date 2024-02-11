package designpattern

import (
	"fmt"
	"sample/designpattern/serverbuilder"
)

func Builder() {
	fmt.Println("=== builder ===")
	// serverbuilder
	srv := serverbuilder.
		NewBuilder("localhost", 8080).
		Timeout(100 * 1000).
		Build()
	srv.Start()
}
