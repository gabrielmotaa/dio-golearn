package main

import "testing"

func TestSumSuccess(t *testing.T) {
	cases := []struct {
		values []int
		result int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{10, 20, 30}, 60},
		{[]int{4, 2}, 6},
	}

	for _, c := range cases {
		result, err := Sum(c.values...)

		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if result != c.result {
			t.Errorf("got: %d, expected: %d", result, c.result)
		}
	}
}

func TestSubtractSuccess(t *testing.T) {
	cases := []struct {
		values []int
		result int
	}{
		{[]int{1, 2, 3}, -6},
		{[]int{-30, -20, -10}, 60},
		{[]int{-100, 10}, 90},
	}

	for _, c := range cases {
		result, err := Subtract(c.values...)

		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if result != c.result {
			t.Errorf("got: %d, expected: %d", result, c.result)
		}
	}
}

func TestMultiplySuccess(t *testing.T) {
	cases := []struct {
		values []int
		result int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{10, 20, 30}, 6_000},
		{[]int{4, -2}, -8},
		{[]int{-2, -2}, 4},
	}

	for _, c := range cases {
		result, err := Multiply(c.values...)

		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if result != c.result {
			t.Errorf("got: %d, expected: %d", result, c.result)
		}
	}
}

func TestDivideSuccess(t *testing.T) {
	cases := []struct {
		values []int
		result int
	}{
		{[]int{20, 10}, 2},
		{[]int{100, 10, 5}, 2},
		{[]int{10}, 10},
		{[]int{10, 3}, 3},
		{[]int{0, 10}, 0},
	}

	for _, c := range cases {
		result, err := Divide(c.values...)

		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if result != c.result {
			t.Errorf("got: %d, expected: %d", result, c.result)
		}
	}
}

func TestDivideError(t *testing.T) {
	cases := []struct {
		values []int
		result int
	}{
		{[]int{20, 0}, 0},
		{[]int{0, 0}, 0},
		{[]int{5, 4, 0, 3}, 0},
	}

	for _, c := range cases {
		result, err := Divide(c.values...)

		if err != ErrCantDivideByZero {
			t.Errorf("expected ErrCantDivideByZero, got: %v", err)
		}

		if result != c.result {
			t.Errorf("got: %d, expected: %d", result, c.result)
		}
	}
}
