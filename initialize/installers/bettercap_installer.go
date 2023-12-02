package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const bettercapCommand = "bettercap"

func BettercapInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkbettercapDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isbettercapInstalled() {
			//fmt.Println("bettercap is already installed.")

			return true

		} else {
			//fmt.Println("Installing bettercap...")
			if err := installbettercap(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("bettercap installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isbettercapInstalled() bool {
	_, err := exec.LookPath(bettercapCommand)
	return err == nil
}

func installbettercap() error {
	cmd := exec.Command("sudo", "apt", "install", "bettercap", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkbettercapDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
