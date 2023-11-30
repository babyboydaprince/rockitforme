package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const niktoCommand = "nikto"

func NiktoInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkniktoDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isniktoInstalled() {
			//fmt.Println("nikto is already installed.")

			return true

		} else {
			//fmt.Println("Installing nikto...")
			if err := installnikto(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("nikto installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isniktoInstalled() bool {
	_, err := exec.LookPath(niktoCommand)
	return err == nil
}

func installnikto() error {
	cmd := exec.Command("sudo", "apt", "install", "nikto", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkniktoDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
