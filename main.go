package main

import (
	"fmt"
	// "reflect"
	// "strconv"
)

func main() {

	// uint8
	// uint16
	// uint32
	// uint64
	// int8
	// int16
	// int32
	// int64
	// float32
	// float64
	// complex64
	// complex128
	// byte
	// rune
	// string
	// bool

	// var x = 1
	// fmt.Println(reflect.TypeOf(x))

	// var welcome string = "hello, world!"
	// welcome := "hello, world!"
	// const PI float64 = 3.14159265359

	// fmt.Println(PI)
	// fmt.Println(welcome)

	// fmt.Println(strconv.ParseInt("15", 10, 64))

	var name string

	fmt.Print("What is your name? ")
	fmt.Scan(&name)

	fmt.Println("Hello, " + name + "!")
}
