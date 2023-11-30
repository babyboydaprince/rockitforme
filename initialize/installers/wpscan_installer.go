package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const wpscanCommand = "wpscan"

func WpscanInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkwpscanDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if iswpscanInstalled() {
			//fmt.Println("wpscan is already installed.")

			return true

		} else {
			//fmt.Println("Installing wpscan...")
			if err := installwpscan(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("wpscan installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func iswpscanInstalled() bool {
	_, err := exec.LookPath(wpscanCommand)
	return err == nil
}

func installwpscan() error {
	cmd := exec.Command("sudo", "apt", "install", "wpscan", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkwpscanDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
