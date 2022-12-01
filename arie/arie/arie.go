package arie

import "fmt"

// internal function
func sayBye() {
	fmt.Println("Bye")
}

// exported function
func SayHello() {
	fmt.Println("Hello")
}
