package main

import "fmt"

// declaration
const pi = 3.1416

// const name := "Ibrahim Bhai"  short declaration hobena

// constant group
const (
	name = "Md. Ibrahim"
	id   = 445
	Uni  = "BUBT"
)

func main() {

	fmt.Println("Pi value is ", pi)
	fmt.Println(name, id, Uni)
}
