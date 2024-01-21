package main

import (
	"fmt"
	"sample/basic"
	"sample/goroutine"
)

func main() {
	fmt.Println("■■■■■■■■■■■■■■■")
	fmt.Println("■■■■ Basic ■■■■")
	fmt.Println("■■■■■■■■■■■■■■■")
	basic.Printtype()
	basic.Convert()
	basic.Operate()
	basic.Format()
	basic.Variable()
	basic.Constant()
	basic.Pointer()
	basic.Array()
	basic.Slice()
	basic.Map()
	basic.Function()
	basic.Struct()
	basic.Interface()
	basic.Statement()
	basic.Defer()
	basic.PanicAndRecover()
	basic.Error()
	fmt.Println("■■■■■■■■■■■■■■■■■■■■")
	fmt.Println("■■■■ Go routine ■■■■")
	fmt.Println("■■■■■■■■■■■■■■■■■■■■")
	goroutine.Goroutine()
	goroutine.Channel()
	//goroutine.Deadlock()
	goroutine.Select()
}
