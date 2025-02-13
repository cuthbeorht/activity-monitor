package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	token := getToken(os.Args)
	fmt.Println("Status: ", token)
}

func getToken(arguments []string) string {
	if len(arguments) != 2 {
		log.Fatal("Invalid number of argumets. Expected 1, got", len(arguments)-1)
	}

	return arguments[1]
}
