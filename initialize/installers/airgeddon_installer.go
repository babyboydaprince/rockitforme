package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const airgeddonCommand = "airgeddon"

func AirgeddonInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkairgeddonDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isairgeddonInstalled() {
			//fmt.Println("airgeddon is already installed.")

			return true

		} else {
			//fmt.Println("Installing airgeddon...")
			if err := installairgeddon(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("airgeddon installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isairgeddonInstalled() bool {
	_, err := exec.LookPath(airgeddonCommand)
	return err == nil
}

func installairgeddon() error {
	cmd := exec.Command("sudo", "apt", "install", "airgeddon", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkairgeddonDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
