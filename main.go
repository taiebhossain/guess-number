package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("!!!Number Guessing Game!!!\n\n")
	fmt.Println("You have three(3) tries ")

	// generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	randomNum := randomizer.Intn(10)
	previousGuess := []int{}

	var guess int
	for try := 1; try <= 3; try++ {
		fmt.Printf("TRIAL: %d\n", try)
		fmt.Println("Guess the number: ")
		fmt.Scanln(&guess)
		previousGuess = append(previousGuess, guess)

		if guess < randomNum {
			fmt.Println("Wrong!!! Your guess is too low")
		} else if guess > randomNum {
			fmt.Println("Wrong!!! Your guess is too high")
		} else {
			fmt.Println("Congratulations!!! You won the game.")
			fmt.Println("Previous guesses: ", previousGuess)
			break
		}

		fmt.Println("Previous guesses: ", previousGuess)

		if try == 3 {
			fmt.Println("!!!GAME OVER!!!")
		}
	}

}
