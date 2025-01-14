package main

import (
	"fmt"
	"os"
)

func main() {
	showGreeting()
	showMenu()

	command := getOption()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 3:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option.")
		os.Exit(-1)
	}
}

// Displays a greeting message to the user
func showGreeting() {
	name := "Coder"
	version := 1.1

	fmt.Printf("Hi, %s!\n", name)
	fmt.Printf("This program is version %.1f\n", version)
}

// Displays the program menu
func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")
}

// Reads and returns the user's selected option
func getOption() int {
	var command int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&command)

	return command
}
