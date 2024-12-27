package main

import (
	"fmt"
)

func main() {
	fmt.Printf("What is your name?: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name,".",)
}
