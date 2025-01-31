// Write a program that takes user input as a string and converts it into an integer and a float

package main

import (
	"fmt"

	"strconv"
)

func convertor() {
	// Task 1: Write a program that takes user input as a string and converts it into an integer and a float
	fmt.Println("Write value to convert into int and float")

	var str string
	fmt.Scan(&str)

	valFloat, _ := strconv.ParseFloat(str, 64)
	valInt := int(valFloat)

	fmt.Println("string to int ", valInt)
	fmt.Printf("string to float %.1f\n", valFloat)
	fmt.Println("int to float ", float64(valInt))
}

func findType() {
	//Implement a simple program that checks the type of a variable dynamically
	fmt.Println("Put the value for checking the value dynamically")

	var val string
	fmt.Scanln(&val)

	if _, err := strconv.Atoi(val); err == nil {
		println("It is of type int")
		return
	}

	if _, err := strconv.ParseFloat(val, 64); err == nil {
		println("It is of type float")
		return
	}

	if _, err := strconv.ParseBool(val); err == nil {
		println("It is of type bool")
		return
	}

	println("It is of type string")

}

// Convert a slice of bytes into a string and vice versa

// In Go, a byte is simply an alias for the uint8 type, which
//
//	can hold values from 0 to 255. These values represent
//	 characters in the ASCII or UTF-8 encoding.
func byteToString() {
	var n int
	fmt.Scan(&n)

	sl := make([]byte, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&sl[i]) //70 72 83
	}

	//byte to slice
	str := string(sl)

	fmt.Printf("%s", str)
}

func stringToByte() {
	var str string
	fmt.Scan(&str)

	byt := []byte(str)
	fmt.Println(byt)
}

const cons1 = "hello-1"

const (
	A = iota
	B = iota
)

// /Write a program that demonstrates the difference between stack and heap allocations using new and make
func diffBetStackAndHeap() {
	//stack allocation
	fmt.Println("Stack allocations")
	a := 5
	fmt.Println(a)

	sli := []int{1, 2, 3, 4, 5}
	fmt.Println(sli)

	fmt.Println("\nHeap allocations using make")
	sli2 := make([]int, 5)
	sli2 = []int{1, 2, 3, 4, 5}
	fmt.Println(sli2)

	fmt.Println("\nHeap allocations using new")
	sli3 := new([]int)
	*sli3 = append(*sli3, 1, 2, 3, 4)

	fmt.Println(*sli3)

}

func practice() {
	// Conditional statements (if, else, switch, fallthrough)
	// a := 10
	// b := 0
	// if a < b {
	// 	fmt.Printf("%d is less than %d\n", a, b)
	// } else {
	// 	fmt.Printf("%d is greater than or equal to %d\n", a, b)
	// }

	//switch case
	// age := 17
	// switch {
	// case age < 18:
	// 	fmt.Println("Age is less than 18")
	// case age < 26:
	// 	fmt.Println("Age is less than 26")
	// default:
	// 	fmt.Println("Age is greater than or equal to 26")
	// }

	// age := 129
	// switch {
	// case age < 18:
	// 	fmt.Println("Age is less than 18")
	// 	fallthrough //if this condition will satisfy then next will definitely occur
	// case age < 26:
	// 	fmt.Println("Age is less than 26")
	// 	fallthrough
	// default:
	// 	fmt.Println("Age is greater than or equal to 26")
	// }

	// Loops: for, range, break, continue
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// n := 10
	// for i := range n {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	//defer statement - LIFO
	fmt.Println("Main fn")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Main-2 fn")

	// x := 5

	// defer fmt.Println("one", x)  // run in last but value of x stored from first without increementing
	// x++
	// defer fmt.Println("two", x)

}

func main() {
	// Task1()

	// Task2()

	//const and iota
	// const cons = 24

	//iota
	// const (
	// 	C = iota
	// 	D
	// )

	const cons = "hello-2"
	// fmt.Println(cons1)
	// fmt.Println(A)
	// fmt.Println(B)
	// fmt.Println(C)
	// fmt.Println(D)

	// byteToString();
	// stringToByte()

	// practice()t
	// diffBetStackAndHeap()
	// convertor()
	findType()

}
