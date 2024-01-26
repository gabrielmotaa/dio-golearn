package main

import (
	"errors"
)

var ErrCantDivideByZero = errors.New("can't divide by zero")

func Sum(i ...int) (int, error) {
	total := 0
	for _, v := range i {
		total += v
	}
	return total, nil
}

func Subtract(i ...int) (int, error) {
	total := 0
	for _, v := range i {
		total = total - v
	}
	return total, nil
}

func Multiply(i ...int) (int, error) {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total, nil
}

func Divide(i ...int) (int, error) {
	total := i[0]
	i = i[1:]
	for _, v := range i {
		if v == 0 {
			return 0, ErrCantDivideByZero
		}
		total = total / v
	}
	return total, nil
}
