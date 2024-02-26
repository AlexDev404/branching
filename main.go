// A hello world program that prints out "Hello, World!" to the console 4 times.
package main

import "fmt"

func decrement(value int) int {
	fmt.Println("Decrementing the value...")
	return value - 1
}

<<<<<<< HEAD
		fmt.Println("The value is:", i)
=======
		fmt.Println("apple")
>>>>>>> master

// The main function
func main() {
	fmt.Println("This program demonstrates functions!")
	// A loop
	// Increments the value passed in by 1
	for i := 0; i < 4; i++ {
		fmt.Println("Hello, World!")
		fmt.Println("apple")
	}
	fmt.Println("Goodbye, World!")
}