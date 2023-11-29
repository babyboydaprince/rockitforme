package sdr_tools

import (
	"fmt"
	"os"
	"os/exec"
	"rockitforme/banner"
	"rockitforme/common"
)

func SDRTools() {

SDRToolsMenuLoop:
	for {

		fmt.Print("\033[H\033[2J") // Clear the console
		banner.BannerSDRTools()

		sdrOptions := common.SingleSelect(
			"\n ----IT'S ALL ABOUT FREQUENCY----\n      ----JUST LIKE SDR----\n\n", []string{
				"GQRX",
				"Go back",
				"Exit",
			})

		if sdrOptions == "" {

			continue
		}

		switch sdrOptions {

		case "GQRX":
			cmd := exec.Command("gnome-terminal", "--", "bash", "-c", "gqrx; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}
			goto SDRToolsMenuLoop

		case "Go back":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			break SDRToolsMenuLoop

		case "Exit":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			fmt.Println("\nExiting...")
			os.Exit(0)

		}

	}
}
