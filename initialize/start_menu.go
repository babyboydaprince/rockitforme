package initialize

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"rockitforme/banner"
	"rockitforme/common"
	"rockitforme/tools/net_tools"
	"rockitforme/tools/nfc_tools"
	"rockitforme/tools/sdr_tools"
	"rockitforme/tools/web_tools"
	"rockitforme/tools/wireless_tools"
	"syscall"
)

var inSubmodule bool

func StartMenu() {

	inSubmodule = false

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh

		fmt.Print("\033[H\033[2J")
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

			fmt.Print("\033[H\033[2J")
			banner.BannerNetTools()

			netOptions := common.SingleSelect("\n  ----TO SHOW YOU THE WAY----\n", []string{
				"Bettercap",
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

			case "Bettercap":

				inSubmodule = false

				net_tools.Bettercap()

				fmt.Print("\033[H\033[2J")
				banner.PrintBanner()

				goto MainMenuLoop

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

				fmt.Print("\033[H\033[2J")
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
				fmt.Print("\033[H\033[2J")
				banner.PrintBanner()

				goto MainMenuLoop

			case "Exit":
				fmt.Print("\033[H\033[2J")
				banner.PrintBanner()

				fmt.Println("\nExiting...")
				os.Exit(0)

			}
			//-----------------------------------------------------------------------------------------------------------------
		case "Wireless":
			inSubmodule = false

			wireless_tools.WirelessTools()

		case "Web":
			inSubmodule = false

			web_tools.WebTools()

		case "SDR":
			inSubmodule = false

			sdr_tools.SDRTools()

		case "NFC & RFID":
			inSubmodule = false

			nfc_tools.NFCTools()

		case "Exit":
			fmt.Print("\033[H\033[2J")
			banner.PrintBanner()

			fmt.Println("\nExiting...")
			os.Exit(0)

		}

	}
}
