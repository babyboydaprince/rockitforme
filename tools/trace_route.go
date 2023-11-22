package tools

import (
	"fmt"
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
			banner.PrintBanner()
			fmt.Printf("Tracing Route... %s\r", spinChars[i])
			i = (i + 1) % len(spinChars)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Trace route scan using nmap
func TraceRoute(host string) {
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
		return
	}

	// Clear the console and print the banner
	fmt.Print("\033[H\033[2J")
	banner.PrintBanner()

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
}
