package foo

import (
	"errors"
	"fmt"
)

func sayBye() error { // private
	fmt.Println("Bye...")
	return nil
}

func SayHi() error { // public
	fmt.Println("Hi!")
	return nil
}

func MakeError() error {
	return errors.New("Error!")
}
