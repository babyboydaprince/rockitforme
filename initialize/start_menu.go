package initialize

import (
	"bufio"
	"fmt"
	"os"
	"rockitforme/banner"
	"rockitforme/common"
	"rockitforme/tools"
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

				fmt.Println("\nSpoofing...")

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
