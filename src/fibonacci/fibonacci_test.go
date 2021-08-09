package main

import "testing"

func Test_fib_one(t *testing.T) {
	actual := fib_one(6)

	if actual != 8 {
		t.Error("fib_one: expected value of 8 at position 6.")
	}
}

func Test_fib_two(t *testing.T) {
	actual := fib_two(6, nil)

	if actual != 8 {
		t.Error("fib_two: expected value of 8 at position 6.")
	}
}

func Test_fib_three(t *testing.T) {
	actual := fib_three(6)

	if actual != 8 {
		t.Error("fib_three: expected value of 8 at position 6.")
	}
}
