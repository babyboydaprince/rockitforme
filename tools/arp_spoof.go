package tools


import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"rockitforme/banner"
	"strings"
	"time"
)

func ArpSpoof(target string, gateway string) {

	fmt.Print("\033[H\033[2J") // Clear the console
	banner.BannerArpSpoof()

	fmt.Print("\n    ----TO KEEP WATCH---\n")

	CheckIPForward()

	fmt.Printf("\nARP Spoofing in progress for Target: %s, Gateway: %s\n", target, gateway)
	fmt.Print("\nPress Ctrl+C to stop and return to the main menu.\n\n")

	cmd := exec.Command("arpspoof", target, "-t", gateway)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command in a goroutine
	go func() {
		if err := cmd.Start(); err != nil {
			fmt.Println("\n\nError starting arpspoof:", err)
			return
		}
		// Wait for the command to finish
		cmd.Wait()
	}()

	// Listen for Ctrl+C signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	// Wait for the termination signal
	<-sigCh

	// Handle the termination gracefully
	fmt.Println("\n\nStopping ARP Spoofing...")
	if err := cmd.Process.Kill(); err != nil {
		fmt.Println("\n\nError stopping arpspoof:", err)
	}

	// Wait for a brief moment to allow the module to finish
	time.Sleep(2 * time.Second)

	// Clear the console and go back to the main menu
	fmt.Print("\033[H\033[2J") // Clear the console
	banner.PrintBanner()
}

func CheckIPForward() {

CheckIPForwardLoop:

	for {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("\n\nNeed to configure IP FORWARDING? (Y/N)\n\nor type BACK to return to menu: ")

		scanner.Scan()

		choiceIPFWD := scanner.Text()

		switch strings.ToLower(choiceIPFWD) {

		case "y":
			ipForwarder(choiceIPFWD)
			break CheckIPForwardLoop

		case "n":
			break CheckIPForwardLoop

		case "back":
			break CheckIPForwardLoop

		default:
			fmt.Print("\nNot an option...")
			time.Sleep(2 * time.Second)

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerArpSpoof()

			goto CheckIPForwardLoop
		}

	}

}

func ipForwarder(choice string) {

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
}
