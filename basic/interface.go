package basic

import (
	"fmt"
)

func Interface() {
	fmt.Println("=== interface ===")
	var a Animal
	a = Dog{Name: "pochi"}
	a.Cry() // bow
	a = Cat{Name: "tama"}
	a.Cry() // meow
	// a = Car{Name: "prius"} // compile error

	var emp interface{}
	emp = 1
	fmt.Println(emp) // 1
	emp = "hello"
	fmt.Println(emp) // hello

}

type Animal interface {
	Cry()
}

// Dog implements Animal
type Dog struct {
	Name string
}

func (d Dog) Cry() {
	fmt.Println("bow")
}

type Cat struct {
	Name string
}

func (c Cat) Cry() {
	fmt.Println("meow")
}

// Car not implements Animal
type Car struct {
	Name string
}
