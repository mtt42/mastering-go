package main

import (
  "fmt"
	"log"
	"os"
)

func main() {
	arguments := os.Args

  fmt.Println(os.TempDir())

	if len(arguments) != 1 {
		log.Fatal("Fatal: Hello World!")
	}
	log.Panic("Panic: Hello World!")
}
