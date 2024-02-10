package library

import (
	"fmt"
	"os"
)

func Os() {
	fmt.Println("=== os ===")
	// args
	fmt.Println("args:", os.Args) // [./main]
	// env
	fmt.Println("env:", os.Environ()[0]) // [PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin]
	// exit
	//os.Exit(1)
	// hostname
	h, _ := os.Hostname()
	fmt.Println("hostname:", h) // hostname
	wd, _ := os.Getwd()
	fmt.Println("getwd:", wd) // /Users/username/go/src/sample
}

func File() {
	// open
	f, err := os.Open("library/os.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// read
	buf := make([]byte, 1024)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}
	// write
	f2, err := os.Create("library/os2.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()
	f2.Write([]byte("package library\n"))
	f2.Write([]byte("func Os2() {\n"))
	f2.Write([]byte("}\n"))
	// remove
	os.Remove("library/os2.go")
}
