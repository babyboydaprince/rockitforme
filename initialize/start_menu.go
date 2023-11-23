package initialize

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"rockitforme/banner"
	"rockitforme/common"
	"rockitforme/tools"
	"syscall"
)

var inSubmodule bool

func StartMenu() {

	inSubmodule = false

	// Listen for Ctrl+C signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Go routine to handle Ctrl+C signal
	go func() {
		<-sigCh

		fmt.Print("\033[H\033[2J") // Clear the console
		banner.PrintBanner()

		if inSubmodule {
			fmt.Print("\n\nReturning to the main menu...\n\nHit ENTER!")

		} else {
			fmt.Print("\n\nExiting...")
			os.Exit(0)
		}
	}()

MainMenuLoop:

	for {

		options := common.SingleSelect(
			"\n     ---THE EVER PRESENT---\n    ----SILENT OBSERVER-----\n\n       Made by:  BraiNiac\n\n", []string{
				"Net tools",
				"Web",
				"Wireless",
				"SDR",
				"NFC & RFID",
				"Exit",
			})

		inSubmodule = false

		if options == "" {
			// If the user presses Enter without choosing an option, stay in the current menu
			continue
		}

		switch options {

		case "Net tools":

			inSubmodule = false

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerNetTools()

			netOptions := common.SingleSelect("\n  ----TO SHOW YOU THE WAY----\n", []string{
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

				inSubmodule = false

				tools.GoChangeMyMac()

			case "Port scanner":

				inSubmodule = false

				tools.PortScan()

			case "Trace ip route":

				inSubmodule = false

				tools.TraceRoute()

			case "Arp spoof":

				inSubmodule = false

				fmt.Print("\033[H\033[2J") // Clear the console
				banner.BannerArpSpoof()

				fmt.Print("\n    ----TO KEEP WATCH---\n")

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("\n\nSet TARGET IP to spoof: ")

				scanner.Scan()

				target := scanner.Text()

				fmt.Print("\nNow set the GATEWAY IP to spoof: ")

				scanner.Scan()

				gateway := scanner.Text()

				tools.ArpSpoof(target, gateway)

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
