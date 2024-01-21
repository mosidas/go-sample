package basic

import (
	"fmt"
)

func Struct() {
	fmt.Println("=== struct ===")
	var a Person
	var b Person = Person{Name: "foo", Age: 20}
	fmt.Println(a, b)       // { 0 0} {foo 20 0}
	fmt.Println(a.String()) // { 0 0}
	b.SetStatus(1)
	fmt.Println(b.String()) // {foo 20 1}

}

type Person struct {
	// public (first letter is upper case)
	Name string
	Age  int
	// private (first letter is lower case)
	status int
}

func (p *Person) SetStatus(status int) {
	p.status = status
}

// method
func (p *Person) String() string {
	return fmt.Sprintf("{%s %d %d}", p.Name, p.Age, p.status)
}
