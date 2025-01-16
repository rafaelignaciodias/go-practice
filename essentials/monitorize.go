package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoring = 3
const delay = 5

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
	fmt.Printf("This program is version %.1f\n\n", version)
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
	fmt.Println()

	return command
}

// Monitors the website list using the HTTP protocol
func startMonitoring() {
	fmt.Println("Monitoring")
	sites := []string{
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/500",
		"https://httpbin.org/status/200",
	}

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}

}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "has loaded successfully.")
	} else {
		fmt.Println("Site:", site, "has not loaded successfully. Status Code:", resp.StatusCode)
	}
}
