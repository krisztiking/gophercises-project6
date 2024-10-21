package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func bytemethod(data []byte, a int) int {
	min, max := 'A', 'Z'
	for _, char := range data {
		ch := rune(char)
		if ch >= min && ch <= max {
			a++
		}
	}
	return a
}

func uppercasemethod(data []byte, a int) int {
	for _, char := range data {
		ch := string(char)
		if ch == strings.ToUpper(ch) {
			a++
		}
	}
	return a
}

func camelcase() {
	filename := "camel.in"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error: ", err)
	}

	answer := 1
	answer2 := answer

	answer = bytemethod(data, answer)
	answer2 = uppercasemethod(data, answer2)
	fmt.Println("The answer is:", answer)
	fmt.Println("The answer2 is:", answer2)
}
