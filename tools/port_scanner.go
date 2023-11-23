package tools

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"rockitforme/banner"
	"strings"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

// spinner function to show a loading spinner
func spinner(done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	spinChars := []string{"|", "/", "-", "\\"}

	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerPortScanner()
			fmt.Print("\n\n")
			fmt.Printf("              Scanning... %s\r", spinChars[i])
			i = (i + 1) % len(spinChars)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// PortScan performs a port scan using nmap
func PortScan() {

	fmt.Print("\033[H\033[2J") // Clear the console
	banner.BannerPortScanner()

	fmt.Print("\n         ----TO SNEAK IN----\n")

	chekInput()

}

func chekInput() {

CheckInputLoop:
	for {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("\n\nInsert URL or IP ADDRESS to scan \n\nor type BACK to return to menu:  ")

		scanner.Scan()

		host := scanner.Text()

		switch strings.ToLower(host) {

		case "":
			fmt.Print("\nNot an option...")
			time.Sleep(2 * time.Second)

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerPortScanner()

			goto CheckInputLoop

		case "back":
			fmt.Print("\nGoing back...")
			time.Sleep(2 * time.Second)

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			break CheckInputLoop

		default:

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerPortScanner()

			nmapPortScan(host)

			break CheckInputLoop
		}

	}
}

func nmapPortScan(host string) {
	cmd := exec.Command("nmap", host)

	// Channel and wait group for synchronization
	done := make(chan struct{})
	var wg sync.WaitGroup

	// Start the spinner animation
	wg.Add(1)
	go spinner(done, &wg)

	out, err := cmd.Output()

	// Stop the spinner when the scan is completed
	close(done)
	wg.Wait()

	if err != nil {

		// If there was any error, print it here
		fmt.Println("\n\nCould not run command:", err)
		return
	}

	// Clear the console and print the banner
	fmt.Print("\033[H\033[2J")
	banner.BannerPortScanner()

	fmt.Print("\n")

	// Create a new table
	t := table.NewWriter()
	t.SetTitle("Nmap Port Scan Results")
	t.AppendHeader(table.Row{"#", "Results"})

	// Split the output by newline character
	lines := strings.Split(string(out), "\n")

	// Append each line to the table
	for i, line := range lines {
		t.AppendRow(table.Row{i + 1, line})
	}

	// Render and print the table
	fmt.Println(t.Render())

	// Optional: Sleep for a short duration to make the output visible before the next iteration
	time.Sleep(1 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user to try another scan or go back to the Net Tools menu
	fmt.Print("\n\nDo you want to perform another scan? (Y/N): ")

	scanner.Scan()

	response := scanner.Text()

	switch strings.ToLower(response) {

	case "y":
		PortScan()

	case "n":
		// Go back to the Net Tools menu
		fmt.Print("\033[H\033[2J")
		banner.PrintBanner()

		return

	default:
		fmt.Print("\nGoing back...")
		time.Sleep(2 * time.Second)

		fmt.Print("\033[H\033[2J") // Clear the console
		banner.PrintBanner()

		return
	}
}
