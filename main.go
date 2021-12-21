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
	var age int
	var month int
	var number int

	fmt.Print("What is your name? ")
	fmt.Scanf("%s", &name)

	fmt.Print("How old are you? ")
	fmt.Scanf("%d", &age)

	if age >= 18 {
		fmt.Println("Hello, " + name + "!")
	} else {
		fmt.Println("Bye, " + name + ".")
	}

	fmt.Print("What month of your brith? ")
	fmt.Scanf("%d", &month)

	switch month {
	case 1:
		fmt.Println("You were born in January.")
	case 2:
		fmt.Println("You were born in February.")
	case 3:
		fmt.Println("You were born in March.")
	case 4:
		fmt.Println("You were born in April.")
	case 5:
		fmt.Println("You were born in May.")
	case 6:
		fmt.Println("You were born in June.")
	case 7:
		fmt.Println("You were born in July.")
	case 8:
		fmt.Println("You were born in August.")
	case 9:
		fmt.Println("You were born in September.")
	case 10:
		fmt.Println("You were born in October.")
	case 11:
		fmt.Println("You were born in November.")
	case 12:
		fmt.Println("You were born in December.")
	default:
		fmt.Println("Bye, " + name + ".")
	}

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &number)

	for i := 1; i <= 10; i++ {
		fmt.Println(number, "x", i, "=", number*i)
	}
}
