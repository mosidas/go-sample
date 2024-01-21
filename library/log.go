package library

import (
	"fmt"
	"log"
	"os"
)

func Log() {
	fmt.Println("=== log ===")
	// Print
	log.Print("Hello World!")           // 2019/01/01 00:00:00 Hello World
	log.Println("Hello World!!")        // 2019/01/01 00:00:00 Hello World
	log.Printf("Hello  %s", "World!!!") // 2019/01/01 00:00:00 Hello World

	// Fatal
	// log.Fatal("Hello World Fatal")    // 2019/01/01 00:00:00 Hello World
	// log.Fatalln("Hello World Fatal!") // 2019/01/01 00:00:00 Hello World
	// log.Fatalf("Hello %s", "World")   // 2019/01/01 00:00:00 Hello World

	// Panic
	// log.Panic("Hello World!")          // 2019/01/01 00:00:00 Hello World
	// log.Panicln("Hello World!!")       // 2019/01/01 00:00:00 Hello World
	// log.Panicf("Hello %s", "World!!!") // 2019/01/01 00:00:00 Hello World

	// SetFlags
	// log.Ldate: the date in the local time zone: 2009/01/23
	// log.Ltime: the time in the local time zone: 01:23:23
	// log.Lmicroseconds: microsecond resolution: 01:23:23.123123.  assumes Ltime.
	// log.Llongfile: full file name and line number: /a/b/c/d.go:23
	// log.Lshortfile: final file name element and line number: d.go:23. overrides Llongfile
	// log.LUTC: if Ldate or Ltime is set, use UTC rather than the local time zone
	// log.Lmsgprefix: move the "prefix" from the beginning of the line to before the message
	// log.LstdFlags: initial values for the standard logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lshortfile | log.LUTC | log.Lmsgprefix | log.LstdFlags)

	// SetPrefix
	log.SetPrefix("Hello World.")
	log.Println("Hello World.") // Hello World 2019/01/01 00:00:00 Hello World

	// SetOutput
	log.SetOutput(os.Stdout)
	log.Println("Hello World..") // Hello World 2019/01/01 00:00:00 Hello World
}
