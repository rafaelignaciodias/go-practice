/*
Website Monitoring Program

This program monitors a list of websites by checking their availability via HTTP requests.
It logs the status of each site (online/offline) to a file and allows users to view the logs
or exit the program through a simple menu interface.
*/

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

// Monitoring defines the number of monitoring cycles
const Monitoring = 3

// Delay defines the delay in seconds between monitoring cycles
const Delay = 5

// DateTimeFormat defines the format for date and time representation
const DateTime = "2006-01-02 15:04:05"

// Command options
const (
	CommandMonitor = 1
	CommandLogs    = 2
	CommandExit    = 3
)

func main() {
	showGreeting()

	for {
		showMenu()
		command := getOption()

		switch command {
		case CommandMonitor:
			startMonitoring()
		case CommandLogs:
			showLogs()
		case CommandExit:
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
	fmt.Println("Starting monitoring...")

	sites := getSiteList()
	if len(sites) == 0 {
		fmt.Println("No sites to monitor. Exiting monitoring.")
		return
	}

	for i := 0; i < Monitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(Delay * time.Second)
		fmt.Println()
	}
	fmt.Println("Monitoring complete.")
}

// Tests a single website's status
func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Printf("Error accessing site %s: %v\n", site, err)
	}

	if resp.StatusCode == 200 {
		fmt.Printf("%s is up. Status Code: %d.\n", site, resp.StatusCode)
		writeSiteLog(site, true)
	} else {
		fmt.Printf("%s is down. Status Code: %d.\n", site, resp.StatusCode)
		writeSiteLog(site, false)
	}
}

// Reads and returns a list of websites from a file
func getSiteList() []string {
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
	}
	var sites []string

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
		fmt.Printf("Failed to open log file: %v\n", err)
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
	fmt.Println("Showing log...")

	fmt.Println(string(file))
}
