package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const mfcukCommand = "mfcuk"

func MfcukInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkmfcukDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if ismfcukInstalled() {
			//fmt.Println("mfcuk is already installed.")

			return true

		} else {
			//fmt.Println("Installing mfcuk...")
			if err := installmfcuk(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("mfcuk installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func ismfcukInstalled() bool {
	_, err := exec.LookPath(mfcukCommand)
	return err == nil
}

func installmfcuk() error {
	cmd := exec.Command("sudo", "apt", "install", "mfcuk", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkmfcukDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
