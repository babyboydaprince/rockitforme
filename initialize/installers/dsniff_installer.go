package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const dsniffCommand = "dsniff"

func DsniffInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkdsniffDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isdsniffInstalled() {
			//fmt.Println("dsniff is already installed.")

			return true

		} else {
			//fmt.Println("Installing dsniff...")
			if err := installdsniff(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("dsniff installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isdsniffInstalled() bool {
	_, err := exec.LookPath(dsniffCommand)
	return err == nil
}

func installdsniff() error {
	cmd := exec.Command("sudo", "apt", "install", "dsniff", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkdsniffDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
