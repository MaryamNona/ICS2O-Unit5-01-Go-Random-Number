// Created by: Maryam Nona
// Createdon: May 2021
//
// This program does basic math

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {
	var randomNumber int
	var userGuess int

rand.seed(time.Now().UnixNano ())
	min := 1
	max := 6
	randomNumber = rand.Intn((max - min + 1) + min)
	userGuess = -1

	fmt.Println("This program chooses a random number between 1-6 ")
	fmt.println()
	// input
	for userGuess != randomNumber {
		fmt.Print("Enter number: ")
		fmt.Scanln(&userGuess)

		// output
		if randomNumber > userGuess {
			fmt.Println("Number too low.")
		}
		if randomNumber < userGuess {
			fmt.Println("Number too high.")
		}
		if randomNumber == userGuess {
		  fmt.Println("Correct number!")
		}
	}
}
