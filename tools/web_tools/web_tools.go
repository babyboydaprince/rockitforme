package web_tools

import (
	"fmt"
	"os"
	"os/exec"
	"rockitforme/banner"
	"rockitforme/common"
)

func WebTools() {

WebToolsMenuLoop:
	for {

		fmt.Print("\033[H\033[2J") // Clear the console
		banner.BannerWebTools()

		webOptions := common.SingleSelect("\n ----ABOUT TO OCCUPY THE WEB----\n", []string{
			"Burpsuite",
			"SQLMAP",
			"NMAP",
			"NIKTO",
			"WPSCAN",
			"Go back",
			"Exit",
		})

		if webOptions == "" {

			continue
		}

		switch webOptions {

		case "Burpsuite":
			cmd := exec.Command("gnome-terminal", "--", "bash", "-c", "burpsuite; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}
			goto WebToolsMenuLoop

		case "SQLMAP":

			cmd := exec.Command("gnome-terminal", "--", "bash", "-c", "sqlmap --help; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}
			goto WebToolsMenuLoop

		case "NMAP":
			cmd := exec.Command("gnome-terminal",
				"--", "bash", "-c", "nmap --help; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}

			goto WebToolsMenuLoop

		case "NIKTO":
			cmd := exec.Command("gnome-terminal",
				"--", "bash", "-c", "nikto -Help; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}

			goto WebToolsMenuLoop

		case "WPSCAN":
			cmd := exec.Command("gnome-terminal",
				"--", "bash", "-c", "wpscan --help; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}

			goto WebToolsMenuLoop

		case "Go back":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			break WebToolsMenuLoop

		case "Exit":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			fmt.Println("\nExiting...")
			os.Exit(0)

		}

	}
}
