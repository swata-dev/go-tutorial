package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Brad"
	var age int32 = 23
	var weight int32 = 55
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "swata"
	// email := "swata@gmail.com"

	name, email, sex := "Swata", "swata@gmail.com", "male"

	fmt.Println(name, age, sex, weight, isCool, email)
	fmt.Printf("%T\n", size)
}

// face
