package main

import (
	"errors"
	"fmt"
	"slices"
)

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return slices.Equal(me.Codes, me2.Codes)
	}
	return false
}

func returnsWrappedMyErr() error {
	return fmt.Errorf("wrapping a MyErr: %w", MyErr{
		Codes: []int{2, 7, 1, 8, 2, 8},
	})
}

func main() {
	err := returnsWrappedMyErr()
	me := MyErr{Codes: []int{2, 7, 1, 8, 2, 8}}
	if errors.Is(err, me) {
		fmt.Println("found it!")
	}
}
