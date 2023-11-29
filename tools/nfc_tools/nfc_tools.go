package nfc_tools

import (
	"fmt"
	"os"
	"os/exec"
	"rockitforme/banner"
	"rockitforme/common"
)

func NFCTools() {

NFCToolsMenuLoop:
	for {

		fmt.Print("\033[H\033[2J") // Clear the console
		banner.BannerNFCTools()

		nfcOptions := common.SingleSelect("\n  ----MAKE CONTACT NO MORE----\n    ----GO CONTACTLESS----\n\n", []string{
			"NFC-TOOLS",
			"MFOC",
			"MFCUK",
			"Go back",
			"Exit",
		})

		if nfcOptions == "" {

			continue
		}

		switch nfcOptions {

		case "NFC-TOOLS":

			cmd := exec.Command("gnome-terminal", "--", "bash", "-c", "nfc-scan-device; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}
			goto NFCToolsMenuLoop

		case "MFOC":
			cmd := exec.Command("gnome-terminal",
				"--", "bash", "-c", "mfoc --help; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}

			goto NFCToolsMenuLoop

		case "MFCUK":
			cmd := exec.Command("gnome-terminal",
				"--", "bash", "-c", "mfcuk --help; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}

			goto NFCToolsMenuLoop

		case "Go back":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			break NFCToolsMenuLoop

		case "Exit":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			fmt.Println("\nExiting...")
			os.Exit(0)

		}

	}
}
