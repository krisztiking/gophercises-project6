package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func rotate(ru rune, va int, key []rune) rune {
	index := -1
	for i, r := range key {
		if ru == r {
			index = i
			break
		}

	}
	if index == -1 {
		return ru
	}
	newindex := index + va
	if newindex <= len(key) {
		return key[newindex]
	} else {
		return key[(newindex-len(key))-1]
	}
}

func caesar() {
	filename := "caesar.in"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error: ", err)
	}
	defer file.Close()

	var line2 string
	var line1, line3 int

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		v := scanner.Text()
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		line1 = num
	}
	if scanner.Scan() {
		line2 = scanner.Text()
	}
	if scanner.Scan() {
		v := scanner.Text()
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		line3 = num
	}

	alphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var new string
	for _, v := range line2 {

		new = new + (string(rotate(v, line3, alphabet)))
	}
	fmt.Println(line1)
	fmt.Println("Normal Sentence:", line2, "\ncoded sentence:", new)

}
