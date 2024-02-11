package designpattern

import (
	"fmt"
	"sample/designpattern/serveroptional"
	"sample/designpattern/serversimple"
)

func FunctionalOptional() {
	fmt.Println("=== functional optional ===")
	// serversimple
	srv := serversimple.New("localhost", 8080)
	srv.Start()

	// serveroptional
	srv2 := serveroptional.New(
		"localhost",
		8080,
		serveroptional.WithTimeout(10),
		serveroptional.WithLogger(nil))
	srv2.Start()
}
