package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1
	// fmt.Println("Hello World!")
	// for i := 0; i < 5; i++ {
	// 	sayMessage("Hello Go!", i)
	// }

	// 2
	// message := "my message"
	// updateMessage(&message)
	// fmt.Println(message)

	// 3
	// divideRes, err := divide(1, 2)
	// if err != nil {
	// 	fmt.Println("something went wrong with division")
	// 	return
	// }
	// fmt.Println(divideRes)

	// // annonymus function
	// func() {
	// 	fmt.Println("hello from annonymus function")
	// }()

	// // passing in values to an annonymous function(e.g. in case of async)
	// for i := 0; i < 5; i++ {
	// 	func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }

	// func as variable
	// var myFunction func()
	// myFunction = func() {
	// 	fmt.Println("this is a portable function!")
	// }
	// myFunction()

	// simulate oop
	myBook := Book{
		id:   1,
		name: "My First Book",
	}

	fmt.Println(myBook.display())
	myBook.changeName("Nolan's Book")
	fmt.Println(myBook.display())
}

//1. if all the params are of same type, just write for the last one
func sayMessage(message string, index int) {
	fmt.Println(message)
	fmt.Println("The value of the index is : ", index)
}

func updateMessage(message *string) {
	*message = "MODIFIED " + *message
}

//2. named variables
func sum(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}

//3. function with error value
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

//4. method ================

type Book struct {
	id   int
	name string
}

func (b Book) display() (message string) {
	return strconv.Itoa(b.id) + " " + b.name
}

func (b *Book) changeName(newName string) {
	b.name = newName
}

//===========================
