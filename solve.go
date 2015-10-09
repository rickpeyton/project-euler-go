package main

import (
	"fmt"
)

func main() {
	fmt.Println(p001())
	fmt.Println(p002())
	multiples_of_99()
}

// Problem 1
func p001() int {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// Problem 2
func p002() int {
	previous := 2
	previous_previous := 1
	sum := 2

	for i := 2; i < 4000000; {
		i = previous + previous_previous
		if i%2 == 0 {
			sum += i
		}
		previous_previous = previous
		previous = i
	}
	return sum
}

// Find all multiples of 99
func multiples_of_99() {
	var multiples []int
	for i := 1; i < 99; i++ {
		if 99%i == 0 {
			multiples = append(multiples, i)
		}
	}
	fmt.Println(multiples)
	return
}
