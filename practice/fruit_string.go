// Code generated by "stringer -type Fruit enum.go"; DO NOT EDIT.

package practice

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Apple-0]
	_ = x[Banana-1]
	_ = x[Cherry-2]
}

const _Fruit_name = "AppleBananaCherry"

var _Fruit_index = [...]uint8{0, 5, 11, 17}

func (i Fruit) String() string {
	if i < 0 || i >= Fruit(len(_Fruit_index)-1) {
		return "Fruit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Fruit_name[_Fruit_index[i]:_Fruit_index[i+1]]
}
