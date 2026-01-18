package main

import (
	"fmt"
)

func main() {
	
	var num1 float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	var num2 float64
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Print("Enter a Operator (+, -, *, /): ")
	var operator string
	fmt.Scan(&operator)

	result := 0.0
	isValid := true

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Divide korte parbana!")
			isValid = false
		} else {
			result = num1 / num2
		}
	default:
		fmt.Println("Invalid Operator!")
		isValid = false
	}

	if isValid {
		fmt.Printf("Answer: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
	}
}