// Recursion
// n! ==> n * (n-1) * (n-2) * (n-3) * ... * 1
// 5! ==> 5 * 4 * 3 * 2 * 1 = 120
// 4! ==> 4 * 3 * 2 * 1 = 24
// 3! ==> 3 * 2 * 1 = 6
// 2! ==> 2 * 1 = 2
// 1! ==> 1
// 0! ==> 1

// A hello world program that prints out "Hello, World!" to the console 4 times.
package main

import (
	"fmt"
	"math"
)

// A function that prints out "Good morning, World!" to the console
func sayGreeting(n string) {
	fmt.Printf("Good morning, %v \n", n)
}

// A function that prints out "Goodbye, World!" to the console
func sayBye(n string) {
	fmt.Printf("Goodbye, %v \n", n)
}

// A function that cycles through a list of names and calls a function for each name
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// A function that calculates the area of a circle
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// A function that decrements the value passed in by 1
func decrement(value int) int {
	fmt.Println("Decrementing the value...")
	return value - 1
}

// A function that increments the value passed in by 1
func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// The main function
func main() {
	fmt.Println("This program demonstrates functions!")
	cycleNames([]string{"Mario", "Luigi", "Peach"}, sayGreeting)
	// A loop
	// Increments the value passed in by 1
	for i := 0; i < 4; i++ {
		sayGreeting("Mario")
		sayGreeting("Luigi")
		// fmt.Println("Hello, World!")
		fmt.Println("apple")
		fmt.Println(factorial(i))
	}
	sayBye("Mario")
	cycleNames([]string{"Mario", "Luigi", "Peach"}, sayBye)
	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %.3f and Circle 2 is %.3f \n", a1, a2)
	// fmt.Println("Goodbye, World!")
}
