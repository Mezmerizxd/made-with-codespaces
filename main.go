package main

import (
	"log"
	"fmt"
)

func main() {
	log.Println("starting...")

	fmt.Println("Type anything below: ")

	var input string
	fmt.Scanln(&input)

	fmt.Println("You typed: '" + input + "'")

	log.Println("exiting...")
}