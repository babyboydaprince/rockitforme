package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const wifiteCommand = "wifite"

func WifiteInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkwifiteDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if iswifiteInstalled() {
			//fmt.Println("wifite is already installed.")

			return true

		} else {
			//fmt.Println("Installing wifite...")
			if err := installwifite(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("wifite installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func iswifiteInstalled() bool {
	_, err := exec.LookPath(wifiteCommand)
	return err == nil
}

func installwifite() error {
	cmd := exec.Command("sudo", "apt", "install", "wifite", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkwifiteDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
