//go:generate stringer -type Fruit enum.go
package practice

type Fruit int

const (
	Apple Fruit = iota
	Banana
	Cherry
)
