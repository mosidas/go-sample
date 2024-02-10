package practice

import (
	"fmt"
)

func Stringer() {
	fmt.Println("=== stringer ===")

	// stringer
	s := &StringOK{
		Id:      1,
		Name:    "Taro",
		Age:     20,
		Hobbies: []string{"Baseball", "Soccer"},
	}
	fmt.Println(s)
}

type StringOK struct {
	Id      int
	Name    string
	Age     int
	Hobbies []string
}

func (s *StringOK) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Age=%d, Hobbies=%v", s.Id, s.Name, s.Age, s.Hobbies)
}
