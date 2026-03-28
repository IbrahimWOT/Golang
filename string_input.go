package main
import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	reader := bufio.NewScanner(os.Stdin)
    fmt.Print("Enter your name : ")
	reader.Scan()
	name := reader.Text()

	fmt.Printf("My name is %s",name)
}
