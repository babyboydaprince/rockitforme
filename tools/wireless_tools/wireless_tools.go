package wireless_tools

import (
	"fmt"
	"os"
	"os/exec"
	"rockitforme/banner"
	"rockitforme/common"
)

func WirelessTools() {

WirelessToolsMenuLoop:
	for {

		fmt.Print("\033[H\033[2J") // Clear the console
		banner.BannerWirelessTools()

		wirelessOptions := common.SingleSelect("\n  ----LOVE IS IN THE AIR----\n", []string{
			"Aircrack-ng Suite",
			"Wifite",
			"Oneshot",
			"Airgeddon",
			"Go back",
			"Exit",
		})

		if wirelessOptions == "" {

			continue
		}

		switch wirelessOptions {

		case "Aircrack-ng Suite":

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.BannerAircrack()

			AircrackOptions := common.SingleSelect("\n  ----RIDERS ON THE STORM----\n", []string{
				"Airbase-ng",
				"Aircrack-ng",
				"Airdecap-ng",
				"Airdecloak-ng",
				//"Airdrop-ng",
				"Aireplay-ng",
				//"Airgraph-ng",
				"Airmon-ng",
				"Airodump-ng",
				"Airolib-ng",
				"Airserv-ng",
				"Airtun-ng",
				"Packetforge-ng",
				"Go back",
				"Exit",
			})

			switch AircrackOptions {

			case "Airbase-ng":

				cmd := exec.Command("gnome-terminal", "--", "bash", "-c", "airbase-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}
				goto WirelessToolsMenuLoop

			case "Aircrack-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "aircrack-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			case "Airdecap-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "airdecap-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			case "Airdecloak-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "airdecloak-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			//case "Airdrop-ng":
			//	cmd := exec.Command("gnome-terminal",
			//		"--", "bash", "-c", "airdrop-ng --help; exec bash")
			//
			//	err := cmd.Run()
			//	if err != nil {
			//		fmt.Println("Error:", err)
			//		return
			//	}
			//	err = cmd.Wait()
			//	if err != nil {
			//		fmt.Println("Error waiting for command:", err)
			//	}
			//
			//	goto WirelessToolsMenuLoop

			case "Aireplay-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "aireplay-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			//case "Airgraph-ng":
			//	cmd := exec.Command("gnome-terminal",
			//		"--", "bash", "-c", "airgraph-ng --help; exec bash")
			//
			//	err := cmd.Run()
			//	if err != nil {
			//		fmt.Println("Error:", err)
			//		return
			//	}
			//	err = cmd.Wait()
			//	if err != nil {
			//		fmt.Println("Error waiting for command:", err)
			//	}
			//
			//	goto WirelessToolsMenuLoop

			case "Airmon-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "airmon-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			case "Airodump-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "airodump-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			case "Airolib-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "airolib-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			case "Airserv-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "airserv-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			case "Airtun-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "airtun-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			case "Packetforge-ng":
				cmd := exec.Command("gnome-terminal",
					"--", "bash", "-c", "packetforge-ng --help; exec bash")

				err := cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				err = cmd.Wait()
				if err != nil {
					fmt.Println("Error waiting for command:", err)
				}

				goto WirelessToolsMenuLoop

			case "Go back":
				fmt.Print("\033[H\033[2J") // Clear the console
				banner.PrintBanner()

				break WirelessToolsMenuLoop

			case "Exit":
				fmt.Print("\033[H\033[2J") // Clear the console
				banner.PrintBanner()

				fmt.Println("\nExiting...")
				os.Exit(0)
			}
			//---------------------------------------------------------------------------------------------------------------
		case "Wifite":
			cmd := exec.Command("gnome-terminal",
				"--", "bash", "-c", "wifite; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}

			goto WirelessToolsMenuLoop

		case "Oneshot":
			// alias oneshot='python /home/user/Documents/Tools/OneShot/oneshot.py'
			// export PATH=$PATH:/usr/local/bin/OneShot
			// export PATH=$PATH:~~/bin
			cmd := exec.Command("gnome-terminal",
				"--", "bash", "-c", "oneshot.py; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}

			goto WirelessToolsMenuLoop

		case "Airgeddon":
			cmd := exec.Command("gnome-terminal",
				"--", "bash", "-c", "airgeddon; exec bash")

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error waiting for command:", err)
			}

			goto WirelessToolsMenuLoop

		case "Go back":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			break WirelessToolsMenuLoop

		case "Exit":
			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()

			fmt.Println("\nExiting...")
			os.Exit(0)

		}

	}
}
