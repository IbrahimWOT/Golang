package main

import "fmt"

func main() {
    age := 18

    if age == 18 {
        fmt.Println("Yes you are eligible")
    } else if age < 18 {
        fmt.Println("You are a child")
    } else {
        fmt.Println("Your age is above 18")
    }
}
