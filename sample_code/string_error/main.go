package main

import (
	"errors"
	"fmt"
)

func doubleEvenErrorsNew(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func doubleEvenFmtErrorf(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

func main() {
	result, err := doubleEvenErrorsNew(1)
	if err != nil {
		fmt.Println(err) // prints "only even numbers are processed"
	}
	fmt.Println(result)

	result, err = doubleEvenFmtErrorf(1)
	if err != nil {
		fmt.Println(err) // prints "1 isn't an even number"
	}
	fmt.Println(result)
}
