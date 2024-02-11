package practice

import (
	"fmt"
)

type Attr struct {
	Name string
	Age  int
}

type AttrEx struct {
	Name string
}

type Teacher struct {
	Attr
	Subject string
}

type Student struct {
	Attr
	AttrEx
	Score int
}

func EmbeddedStruct() {
	fmt.Println("=== Embedded struct ===")
	teacher := Teacher{
		Attr:    Attr{Name: "Taro", Age: 30},
		Subject: "Math",
	}
	fmt.Println("teacher.name:", teacher.Name)
	fmt.Println("teacher.age:", teacher.Age)
	student := Student{
		Attr:   Attr{Name: "Hanako", Age: 20},
		AttrEx: AttrEx{Name: "HanakoEx"},
		Score:  90,
	}
	fmt.Println("student.name:", student.Attr.Name)
	fmt.Println("student.nameex:", student.AttrEx.Name)
	fmt.Println("student.age:", student.Age)
}
