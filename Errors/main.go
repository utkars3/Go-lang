package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("resource not found")

// var ErrNotFound2 = errors.New("resource not found-2")

func getData() error {
	return ErrNotFound
}

func fetchData(id int) error {
	return fmt.Errorf("failed to fetch data for ID %d: %w", id, ErrNotFound)
}

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func doSomething() error {
	return &MyError{Code: 400, Msg: "Bad Request"}
}

func riskyFunction() {
	panic("Something went wrong!")
}

func safeFunction() {

	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	panic("Something went wrong!")
	fmt.Println("This will not execute")
}

func main() {
	// fmt.Println(getData())
	// fmt.Println(fetchData(23))

	// err := fetchData(123)

	// if errors.Is(err, ErrNotFound) {
	// 	fmt.Println("Handle: Resource not found")
	// } else {
	// 	fmt.Println("Handle: Other error")
	// }

	// err := doSomething()

	// var myErr *MyError
	// if errors.As(err, &myErr) {
	//     fmt.Println("Custom error occurred:", myErr.Code, myErr.Msg)
	// }

	// fmt.Println("Before panic")
	// riskyFunction()            // Program crashes here
	// fmt.Println("After panic") // This will not execute

	fmt.Println("Before calling safeFunction")
	safeFunction()
	fmt.Println("After calling safeFunction") // Execution continues

}
