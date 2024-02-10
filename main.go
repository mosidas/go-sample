package main

import (
	"fmt"
	"sample/basic"
	"sample/build_constraint"
	"sample/goroutine"
	"sample/library"
	"sample/practice"
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
	fmt.Println("■■■■■■■■■■■■■■■■■")
	fmt.Println("■■■■ Library ■■■■")
	fmt.Println("■■■■■■■■■■■■■■■■■")
	library.Log()
	library.Strconv()
	library.Time()
	library.Duration()
	library.ConvertTime()
	library.Os()
	library.Path()
	library.Filepath()
	library.Context()
	fmt.Println("■■■■■■■■■■■■■■■■■■")
	fmt.Println("■■■■ Practice ■■■■")
	fmt.Println("■■■■■■■■■■■■■■■■■■")
	practice.Builtin()
	practice.Stringer()
	fmt.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■")
	fmt.Println("■■■■ Build constraint ■■■■")
	fmt.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■")
	buildconstraint.Build()
	buildconstraint.Name() // tag:cat -> cat, tag:!cat -> dog
}
