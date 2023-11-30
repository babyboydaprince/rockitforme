package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const mfocCommand = "mfoc"

func MfocInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkmfocDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if ismfocInstalled() {
			//fmt.Println("mfoc is already installed.")

			return true

		} else {
			//fmt.Println("Installing mfoc...")
			if err := installmfoc(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("mfoc installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func ismfocInstalled() bool {
	_, err := exec.LookPath(mfocCommand)
	return err == nil
}

func installmfoc() error {
	cmd := exec.Command("sudo", "apt", "install", "mfoc", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkmfocDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
