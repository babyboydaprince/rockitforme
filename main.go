package main

import (
	"fmt"
	"os"
	"rockitforme/banner"
	"rockitforme/initialize"
	"runtime"
)

func main() {

	if !isRootOrAdmin() {
		banner.PrintBanner()
		fmt.Println("\n\nIt seems like you are not all that ready to rock...\n\n" +
			"Must be run as root (Unix) or administrator (Windows). \n\nExiting...")
		os.Exit(1)
	}

	initialize.CheckForDependencies()

	banner.PrintBanner()
	initialize.StartMenu()
}

func isRootOrAdmin() bool {
	switch runtime.GOOS {
	case "windows":
		return isAdminWindows()
	case "linux", "darwin":
		return isRootUnix()
	default:
		return false
	}
}

func isAdminWindows() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func isRootUnix() bool {
	return os.Geteuid() == 0
}
