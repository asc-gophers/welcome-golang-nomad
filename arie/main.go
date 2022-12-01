package main

import (
	"fmt"
	"strings"

	"github.com/Jeongah-Shin/learngo/arie"
)

func multiply(a int, b int) int {
	return a * b
}

// [arie jerry jack ryan]
func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return
func lenAndUpperNaked(name string) (length int, uppercase string) {
	// run after return statement
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	fmt.Println("Noooooomad")
	arie.SayHello()
	fmt.Println(multiply(3, 125))

	fmt.Println(lenAndUpper("arie"))
	totalLen, upperName := lenAndUpper("arie")
	fmt.Println(totalLen, upperName)

	const name1 string = "arie"
	// name1 = "nico"
	const isArie bool = true

	var name2 string = "arie"
	name2 = "nico"
	fmt.Println(name2)

	// Only available in function and type of variable
	name3 := "arie" // Go will guess what the type of the data is.
	fmt.Println(name3)

	repeatMe("arie", "jerry", "jack", "ryan")
	fmt.Println(lenAndUpperNaked("jerry"))

}
