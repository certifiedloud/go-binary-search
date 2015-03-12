package main

import "fmt"

func main() {
	// initialize variables
	low := 0
	high := 100 
	attempts := 0
	keep_guessing := true
	// print vague instructions
	fmt.Println("Pick a number between ", low, " and ", high)
	
	// loop while keep_guessing is true
	for keep_guessing == true {
		guess := (high + low)/2
		var response int
		// increment the number of attempts
		attempts++
		// print prompt after each guess
		fmt.Println("My guess is", guess)
		fmt.Println("1=TOO HI")
		fmt.Println("2=TOO LOW")
		fmt.Println("3=RIGHT ON")
		// take in repsonse and capture error (if any)
		_, err := fmt.Scan(&response)
		if err != nil {
			fmt.Println("there was an error: ", err)
		}
		fmt.Println(response)
		// check the response and adjust the next guess accordingly
		if response == 1 {
			fmt.Println("Too high, eh?")
			high = guess -1
		} else if response == 2 {
			fmt.Println("I guessed low, huh?")
			low = guess + 1
		} else if response == 3 {
			fmt.Println("That was easy. Tries: ", attempts)
			// change variable so loop stops
			keep_guessing = false
		}
	}
	// exit the loop
}
