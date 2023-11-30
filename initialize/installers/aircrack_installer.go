package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const aircrackNgCommand = "aircrack-ng"

func AircrackInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkAircrackDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isAircrackNgInstalled() {
			//fmt.Println("Aircrack-ng is already installed.")

			return true

		} else {
			//fmt.Println("Installing Aircrack-ng...")
			if err := installAircrackNg(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("Aircrack-ng installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isAircrackNgInstalled() bool {
	_, err := exec.LookPath(aircrackNgCommand)
	return err == nil
}

func installAircrackNg() error {
	cmd := exec.Command("sudo", "apt", "install", "aircrack-ng", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkAircrackDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
