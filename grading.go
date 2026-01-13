package main

import "fmt"

func main() {
    var a int
	fmt.Print("Enter a number :")
    fmt.Scanf("%d", &a)  // Correct way to take integer input

    if a<=100 && a >= 80 {
        fmt.Println("A+")
    } else if a<80 && a >= 70 {
        fmt.Println("A")
    } else if a<70 && a >= 60 {
        fmt.Println("B")
    } else {
        fmt.Println("Low")
    }
}
