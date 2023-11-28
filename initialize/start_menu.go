package initialize

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"rockitforme/banner"
	"rockitforme/common"
	"rockitforme/tools/net_tools"
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
				"Wireless",
				"Web",
				"SDR",
				"NFC & RFID",
				"Exit",
			})

		inSubmodule = false

		if options == "" {

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

				continue
			}

			switch netOptions {

			case "MAC Changer":

				inSubmodule = false

				net_tools.GoChangeMyMac()

			case "Port scanner":

				inSubmodule = false

				net_tools.PortScan()

			case "Trace ip route":

				inSubmodule = false

				net_tools.TraceRoute()

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

				net_tools.ArpSpoof(target, gateway)

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

		case "Wireless":

			inSubmodule = false

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerWirelessTools()

			wirelessOptions := common.SingleSelect("\n  ----LOVE IS IN THE AIR----\n", []string{
				"Aircrack-ng",
				"Aireplay-ng",
				"Airmon-ng",
				"Airodump-ng",
				"Packetforge-ng",
				"Wifite",
				"Oneshot",
				"Go back",
				"Exit",
			})

			if wirelessOptions == "" {

				continue
			}

			switch wirelessOptions {

			case "Aircrack-ng":

				inSubmodule = false

			case "Aireplay-ng":

				inSubmodule = false

			case "Airmon-ng":

				inSubmodule = false

			case "Airodump-ng":

				inSubmodule = false

			case "Packetforge-ng":

				inSubmodule = false

			case "Wifite":

				inSubmodule = false

			case "Oneshot":

				inSubmodule = false

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
