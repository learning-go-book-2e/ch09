package main

import (
	"errors"
	"fmt"
)

type MyErr struct {
	Codes []int
}

func (me MyErr) CodeVals() []int {
	return me.Codes
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func AFunctionThatReturnsAnError() error {
	return MyErr{Codes: []int{1, 1, 2, 3, 5, 8}}
}

func main() {
	err := AFunctionThatReturnsAnError()
	var myErr MyErr
	if errors.As(err, &myErr) {
		fmt.Println(myErr.Codes)
	}

	var coder interface {
		CodeVals() []int
	}
	if errors.As(err, &coder) {
		fmt.Println(coder.CodeVals())
	}
}
