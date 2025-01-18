package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5
const DateTime = "2006-01-02 15:04:05"

func main() {
	showGreeting()

	for {
		showMenu()
		command := getOption()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
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

	// sites := []string{
	// 	"https://httpbin.org/status/404",
	// 	"https://httpbin.org/status/200",
	// 	"https://httpbin.org/status/500",
	// 	"https://httpbin.org/status/200",
	// }

	sites := getSiteList()

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
}

// Tests a single website's status
func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "has loaded successfully.")
		writeSiteLog(site, true)
	} else {
		fmt.Println("Site:", site, "has not loaded successfully. Status Code:", resp.StatusCode)
		writeSiteLog(site, false)
	}
}

// Reads and returns a list of websites from a file
func getSiteList() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		row = strings.TrimSpace(row)
		sites = append(sites, row)
	}
	defer file.Close()

	return sites
}

// Writes the log for a site's status
func writeSiteLog(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	file.WriteString(time.Now().Format(DateTime) + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
}

// Displays the log file content
func showLogs() {
	file, err := os.ReadFile("logs.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return
	}

	fmt.Println(string(file))
}
