package main

import (
	"fmt"
	"os"
	"rockitforme/banner"
	"rockitforme/initialize"
	"rockitforme/utils"
)

func main() {

	if !utils.IsRootOrAdmin() {
		banner.PrintBanner()
		fmt.Println("\n\nIt seems like you are not all that ready to rock...\n\n" +
			"Must be run as root (Unix) or administrator (Windows). \n\nExiting...")
		os.Exit(1)
	}
	OpSystem := utils.GetLinuxDistro()
	initialize.CheckForDependencies(OpSystem)
	banner.PrintBanner()
	initialize.StartMenu()
}
