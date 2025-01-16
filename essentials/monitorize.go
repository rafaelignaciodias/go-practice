package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showGreeting()

	for {
		showMenu()
		command := getOption()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			startMonitoring()
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option.")
			os.Exit(-1)
		}
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

// Monitores the website list using the HTTP protocol
func startMonitoring() {
	site := "https://httpbin.org/status/404"
	resp, _ := http.Get(site)
	
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "has loaded successfully.")
	} else {
		fmt.Println("Site:", site, "has not loaded successfully. Status Code:", resp.StatusCode)
	}
}
