package initialize

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"rockitforme/banner"
	"rockitforme/common"
	"rockitforme/tools"
	"sync"
	"syscall"
	"time"
)

func StartMenu() {

MainMenuLoop:
	for {
		options := common.SingleSelect("\n       Made by:  BraiNiac\n", []string{
			"Net tools",
			"Web",
			"Wireless",
			"SDR",
			"NFC & RFID",
			"Exit",
		})

		if options == "" {
			// If the user presses Enter without choosing an option, stay in the current menu
			continue
		}

		switch options {

		case "Net tools":

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			netOptions := common.SingleSelect("\n       ----NET TOOLS----\n", []string{
				"MAC Changer",
				"Port scanner",
				"Trace ip route",
				"Arp spoof",
				"Go back",
				"Exit",
			})

			if netOptions == "" {
				// If the user presses Enter without choosing an option, stay in the current menu
				continue
			}

			switch netOptions {

			case "MAC Changer":
				tools.GoChangeMyMac()

			case "Port scanner":
				fmt.Print("\033[H\033[2J") // Clear the console
				banner.PrintBanner()

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("\nInsert URL or IP ADDRESS to scan: ")

				scanner.Scan()

				input := scanner.Text()

				tools.PortScan(input)

			case "Trace ip route":
				fmt.Print("\033[H\033[2J") // Clear the console
				banner.PrintBanner()

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("\nInsert URL or IP ADDRESS to trace the route ")

				scanner.Scan()

				input := scanner.Text()

				tools.TraceRoute(input)

			case "Arp spoof":
				fmt.Print("\033[H\033[2J") // Clear the console
				banner.PrintBanner()

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("\nSet TARGET IP to spoof: ")

				scanner.Scan()

				target := scanner.Text()

				fmt.Print("\nNow set the GATEWAY IP to spoof: ")

				scanner.Scan()

				gateway := scanner.Text()

				// Channel to receive Ctrl+C signal
				sigChan := make(chan os.Signal, 1)
				signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

				// Channel to notify the ARP spoofing goroutine to stop
				stopChan := make(chan struct{})

				// Wait group for synchronization
				var wg sync.WaitGroup

				// Start ARP spoofing command in the background
				wg.Add(1)
				go func() {
					defer wg.Done()
					tools.ArpSpoof(target, gateway, stopChan)
				}()

				// Wait for Ctrl+C signal or ARP spoofing completion
				select {
				case <-sigChan:
					// If Ctrl+C is pressed, interrupt ARP spoofing
					fmt.Println("\nARP Spoofing interrupted. Returning to the main menu.")
					stopChan <- struct{}{} // Notify the ARP spoofing to stop
				case <-time.After(5 * time.Second): // Adjust the duration based on your needs
					// After a certain duration, assume the ARP spoofing command has completed
					fmt.Println("\nARP Spoofing completed!")
					time.Sleep(2 * time.Second)

				}

				// Wait for the ARP spoofing goroutine to finish
				wg.Wait()

			case "Go back":
				fmt.Print("\033[H\033[2J") // Clear the console
				banner.PrintBanner()

				goto MainMenuLoop

			case "Exit":
				fmt.Print("\033[H\033[2J") // Clear the console
				banner.PrintBanner()

				fmt.Println("\nExiting...")
				os.Exit(0)

			}

		case "Web":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			fmt.Println("Web boy")

		case "Wireless":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			fmt.Println("Wifi boy")

		case "SDR":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			fmt.Println("SDR boy")

		case "NFC & RFID":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			fmt.Println("NFC boy")

		case "Exit":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			fmt.Println("\nExiting...")
			os.Exit(0)

		}

	}
}
