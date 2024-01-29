package library

import (
	"encoding/json"
	"fmt"
	"time"
)

func Time() {
	fmt.Println("=== time ===")
	t := time.Now()
	fmt.Printf("%T %v\n", t, t) // time.Time 2019-01-01 00:00:00 +0000 UTC
	// format(yyyy-mm-dd hh:mm:ss)
	fmt.Println(t.Format("2006-01-02 15:04:05"))        // 2019-01-01 00:00:00
	fmt.Println(t.Format("2006-01-02"))                 // 2019-01-01
	fmt.Println(t.Format("15:04:05"))                   // 00:00:00
	fmt.Println(t.Format("2006-01-02 15:04:05.000"))    // 2019-01-01 00:00:00.000
	fmt.Println(t.Format("2006-01-02 15:04:05.000000")) // 2019-01-01 00:00:00.000000
	fmt.Println(t.Format("2006-01-02 Mon"))             // 2019-01-01 Tue
	fmt.Println(t.Format("2006-01-02 Monday"))          // 2019-01-01 Tuesday
	// format(time.*)
	fmt.Println(t.Format(time.RFC3339))     // 2019-01-01T00:00:00Z
	fmt.Println(t.Format(time.RFC3339Nano)) // 2019-01-01T00:00:00Z
	fmt.Println(t.Format(time.RFC822))      // 01 Jan 19 00:00 UTC
	fmt.Println(t.Format(time.RFC822Z))     // 01 Jan 19 00:00 +0000
	fmt.Println(t.Format(time.RFC850))      // Tuesday, 01-Jan-19 00:00:00 UTC

	// Date
	t2 := time.Date(2019, 1, 2, 3, 4, 5, 6, time.UTC)
	fmt.Printf("%T %v\n", t2, t2) // time.Time 2019-01-02 03:04:05.000000006 +0000 UTC
	fmt.Println(t2.Year())        // 2019
	fmt.Println(t2.Month())       // January
	fmt.Println(t2.Day())         // 2
	fmt.Println(t2.Hour())        // 3
	fmt.Println(t2.Minute())      // 4
	fmt.Println(t2.Second())      // 5
	fmt.Println(t2.Nanosecond())  // 6
	fmt.Println(t2.Weekday())     // Wednesday
	fmt.Println(t2.Zone())        // UTC 0
}

func Duration() {
	fmt.Println("=== duration ===")
	fmt.Println("time.Second:", time.Second)           // 1s
	fmt.Println("time.Minute:", time.Minute)           // 1m0s
	fmt.Println("time.Hour:", time.Hour)               // 1h0m0s
	fmt.Println("time.Millisecond:", time.Millisecond) // 1ms
	fmt.Println("time.Microsecond:", time.Microsecond) // 1µs
	fmt.Println("time.Nanosecond:", time.Nanosecond)   // 1ns

	// add
	t := time.Now()
	fmt.Println("t :", t) // 2019-01-01 00:00:00 +0000 UTC
	t2 := time.Now().Add(time.Hour * 2)
	fmt.Println("t2:", t2) // 2019-01-01 01:00:00 +0000 UTC
	// sub
	fmt.Println("sub:", t.Sub(t2)) // -2h0m0s
	// before, after, equal
	fmt.Println("t < t2:", t.Before(t2)) // true
	fmt.Println("t > t2:", t.After(t2))  // false
	fmt.Println("t == t2:", t.Equal(t2)) // false
	// sleep
	fmt.Println("start sleep")
	time.Sleep(time.Second * 2)
	fmt.Println("end sleep")
}

func ConvertTime() {
	fmt.Println("=== convert ===")
	// time -> timestamp
	t := time.Now()
	fmt.Println("t:", t)                       // 2019-01-01 00:00:00 +0000 UTC
	fmt.Println("t.Unix():", t.Unix())         // 1546300800
	fmt.Println("t.UnixNano():", t.UnixNano()) // 1546300800000000000
	// timestamp -> time
	t2 := time.Unix(1546300800, 0)
	fmt.Println("t2:", t2) // 2019-01-01 00:00:00 +0000 UTC
	t3 := time.Unix(0, 1546300800000000000)
	fmt.Println("t3:", t3) // 2019-01-01 00:00:00 +0000 UTC

	// string -> time
	t4, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
	fmt.Println("t4:", t4) // 2019-01-01 00:00:00 +0000 UTC

	// time -> json
	type tJson struct {
		Timestamp time.Time `json:"timestamp"`
	}
	tj, _ := json.Marshal(tJson{Timestamp: t})
	fmt.Println("tj:", string(tj)) // {"timestamp":"2019-01-01T00:00:00Z"}

}
