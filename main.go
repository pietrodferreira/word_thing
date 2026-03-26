package main

import (
	"flag"
	"fmt"
)

func main() {
	wordFlag := flag.String("word", "", "The word to")
	flag.Parse()
	makeText(*wordFlag)
}

func makeText(message string) {
	chars := []byte(message)
	messageLenght := len(message)	
	for i := range messageLenght {
		fmt.Printf("%s\n", chars[:len(chars)-i])	
	}

	for i := 1; i < messageLenght; i++ {
		fmt.Printf("%s\n", chars[:i+1])
	}

}
