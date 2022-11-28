package main

import (
	"fmt"
	"strings"
	"welcome-golang-nomad/jerry/1_Theory/foo"
)

func lenAndUpper(s string) (length int, uppercase string) {
	defer fmt.Println("Finishing with lenAndUpper()..")
	return len(s), strings.ToUpper(s)
}

func main() {
	const const_name string = "Jerry_constant"
	fmt.Println(const_name)

	var name_type_definition string = "Jerry_type_definition"
	fmt.Println(name_type_definition)

	name_type_assertion := "Jerry_type_assertion"
	fmt.Println(name_type_assertion)

	bool_type_assertion := false
	fmt.Println("Bool type assertion:", bool_type_assertion)

	err := foo.SayHi()
	if err == nil {
		fmt.Println("No error")
	}

	// foo.sayBye()

	err = foo.MakeError()
	if err != nil {
		fmt.Println(err.Error())
	}

	length, uppercase := lenAndUpper("jerry")
	fmt.Println(length, uppercase)
}
