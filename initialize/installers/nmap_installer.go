package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const nmapCommand = "nmap"

func NmapInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkNmapDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isNmapInstalled() {
			//fmt.Println("nmap is already installed.")

			return true

		} else {
			//fmt.Println("Installing nmap...")
			if err := installNmap(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("nmap installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isNmapInstalled() bool {
	_, err := exec.LookPath(nmapCommand)
	return err == nil
}

func installNmap() error {
	cmd := exec.Command("sudo", "apt", "install", "nmap", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkNmapDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
