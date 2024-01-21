package basic

import (
	"fmt"
)

func Error() {
	fmt.Println("=== error ===")
	// standard
	if err := raiseError(); err != nil {
		fmt.Println(err)
	}
	// raiseError start
	// error!

	// declare
	err := raiseError()
	if err != nil {
		fmt.Println(err)
	}
	// raiseError start
	// error!
}

func raiseError() error {
	fmt.Println("raiseError start")
	return fmt.Errorf("error: %s", "this is error!")
}
