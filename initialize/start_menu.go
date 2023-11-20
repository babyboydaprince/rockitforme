package initialize

import (
	"fmt"
	"os"
	"rockitforme/common"
	"rockitforme/tools"
)

func StartMenu() {

	options := []string{
		"Net tools",
		"Wifi tools",
		"Exit",
	}

	choice := common.SingleSelect("\nChoose the tool set to use:\n", options)

	if choice == "1" || choice == "Net tools" {
		tools.GoChangeMyMac()
	} else if choice == "2" || choice == "Wifi tools" {
		fmt.Println("Wifi boy")
	} else if choice == "3" || choice == "Exit" || choice == "exit" {
		fmt.Println("\nExiting...")
		os.Exit(0)
	} else {
		fmt.Println("\nInvalid menu choice:", choice)
		fmt.Println(" ")
		// You can add further validation or handling here
	}
	return
}
