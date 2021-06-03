// Created by: Maryam Nona
// Created on: May 2021
//
// This program does basic math

package main

import (
  "fmt"
  "math/rand"
	"time"
)

func main() {
	var randomNumber int
	var userGuess int

	// input
	fmt.Println("This program chooses a random number ")
	fmt.Println()
	fmt.Print("Enter number: ")
	fmt.Scanln(&randomNumber)

	// output
		fmt.Println()
		if randomNumber > userGuess {
    fmt.Println("Number too high.")
}
if randomNumber < userGuess {
	fmt.Println("Number too low.")
}
if randomNumber == userGuess {
	fmt.Println("Correct number!")
}
rand.Seed(time.Now().UnixNano())
  min := 1
	max := 6
fmt.Println(rand.Intn((max - min + 1) + min))
}
