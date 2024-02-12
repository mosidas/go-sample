package practice

import (
	"fmt"
)

type Attr struct {
	Name string
	Age  int
}

func (a *Attr) String() string {
	return fmt.Sprintf("Name=%s, Age=%d", a.Name, a.Age)
}

type AttrEx struct {
	Name string
}

func (a *AttrEx) String() string {
	return fmt.Sprintf("NameEx=%s", a.Name)
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
	fmt.Println(teacher.String()) // suger of teacher.Attr.String()
	student := Student{
		Attr:   Attr{Name: "Hanako", Age: 20},
		AttrEx: AttrEx{Name: "HanakoEx"},
		Score:  90,
	}
	fmt.Println("student.name:", student.Attr.Name)
	fmt.Println("student.nameex:", student.AttrEx.Name)
	fmt.Println("student.age:", student.Age)
	fmt.Println(student.Attr.String())
}
