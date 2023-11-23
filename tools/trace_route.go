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
func spinnerAnim(done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	spinChars := []string{"|", "/", "-", "\\"}

	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerTraceRoute()
			fmt.Print("\n\n")
			fmt.Printf("        Tracing Route... %s\r", spinChars[i])
			i = (i + 1) % len(spinChars)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Trace route scan using nmap
func TraceRoute() {

	fmt.Print("\033[H\033[2J") // Clear the console
	banner.BannerTraceRoute()

	fmt.Print("\n ----TO WALK YOU THE PATH----\n")

	chekTraceInput()

}

func chekTraceInput() {

CheckInputLoop:
	for {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("\n\nInsert URL or IP ADDRESS to trace the route \n\nor type BACK to return to menu:  ")

		scanner.Scan()

		input := scanner.Text()

		switch strings.ToLower(input) {

		case "":
			fmt.Print("\nNot an option...")
			time.Sleep(2 * time.Second)

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerTraceRoute()

			goto CheckInputLoop

		case "back":
			fmt.Print("\nGoing back...")
			time.Sleep(2 * time.Second)

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			break CheckInputLoop

		default:

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerTraceRoute()

			nmapTraceScan(input)

			break CheckInputLoop
		}

	}
}

func nmapTraceScan(host string) {

	cmd := exec.Command("nmap", host, "--traceroute")

	// Channel and wait group for synchronization
	done := make(chan struct{})
	var wg sync.WaitGroup

	// Start the spinner animation
	wg.Add(1)
	go spinnerAnim(done, &wg)

	out, err := cmd.Output()

	// Stop the spinner when the scan is completed
	close(done)
	wg.Wait()

	if err != nil {
		// If there was any error, print it here
		fmt.Println("\n\nCould not run command:", err)
		fmt.Print("\nMust be run as root.\n\n")
		return
	}

	// Clear the console and print the banner
	fmt.Print("\033[H\033[2J")
	banner.BannerTraceRoute()

	// Create a new table
	t := table.NewWriter()
	t.SetTitle("Nmap trace route scan Results")
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
		TraceRoute()

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
