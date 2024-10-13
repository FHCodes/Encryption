package main

import (
	"bufio"              // Package for buffered I/O
	"criptografia/model" // Importing the model package for encryption and decryption functions
	"fmt"                // Package for formatted I/O
	"os"                 // Package for OS-related functionalities
	"strings"            // Package for string manipulation
)

// Function to encrypt or decrypt a phrase based on user input
func encryptDecrypt(phrase string) {
	var option int8 // Variable to store the user's choice (1 for encryption, 2 for decryption)

	reader := bufio.NewReader(os.Stdin) // Creating a new buffered reader for standard input

	// Prompting the user for an action
	fmt.Println("[1] - Encrypt the phrase")
	fmt.Println("[2] - Decrypt the phrase")
	fmt.Scan(&option)       // Reading the user's choice
	reader.ReadString('\n') // Consuming the newline character

	// Check the user's option and call the appropriate function
	if option == 1 {
		model.Encrypt(&phrase)                    // Encrypting the phrase
		fmt.Println("Encrypted phrase: ", phrase) // Displaying the encrypted phrase
	} else if option == 2 {
		model.Decrypt(&phrase)                    // Decrypting the phrase
		fmt.Println("Decrypted phrase: ", phrase) // Displaying the decrypted phrase
	} else {
		// Handle invalid option
		fmt.Println("Invalid option. Try again.")
		os.Exit(1) // Exit the program with an error status
	}
}

// Function to read a phrase from the user
func readInput() string {
	reader := bufio.NewReader(os.Stdin) // Creating a new buffered reader for standard input

	fmt.Println("Enter the phrase: ")            // Prompting the user to enter a phrase
	inputPhrase, _ := reader.ReadString('\n')    // Reading the phrase until newline
	inputPhrase = strings.TrimSpace(inputPhrase) // Removing any leading/trailing whitespace
	return inputPhrase                           // Returning the cleaned phrase
}

// Function to finalize the program based on user input
func finalizeProgram() {
	var exit string = "" // Variable to store the user's decision to continue or exit

	// Loop until the user provides a valid input to exit or continue
	for exit != "n" && exit != "y" {
		reader := bufio.NewReader(os.Stdin) // Creating a new buffered reader for standard input

		// Prompting the user for their choice
		fmt.Println("[y] - To continue")
		fmt.Println("[n] - To exit the program")
		fmt.Scan(&exit) // Reading the user's input

		reader.ReadString('\n') // Consuming the newline character
		if exit == "n" {
			os.Exit(1) // Exit the program if the user chooses 'n'
		} else if exit == "y" {
			fmt.Println("Continuing...") // Informing the user that the program will continue
		} else {
			// Handle invalid input
			fmt.Println("Invalid input. Please type 'y' or 'n'.")
		}
	}
}

// Main function to run the program
func main() {
	var phrase string // Variable to store the user's phrase
	for {
		phrase = readInput() // Read the input phrase from the user

		encryptDecrypt(phrase) // Call the function to encrypt or decrypt the phrase
		finalizeProgram()      // Call the function to finalize the program
	}
}
