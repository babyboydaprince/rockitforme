package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const burpsuiteCommand = "burpsuite"

func BurpsuiteInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkburpsuiteDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isburpsuiteInstalled() {
			//fmt.Println("burpsuite is already installed.")

			return true

		} else {
			//fmt.Println("Installing burpsuite...")
			if err := installburpsuite(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("burpsuite installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isburpsuiteInstalled() bool {
	_, err := exec.LookPath(burpsuiteCommand)
	return err == nil
}

func installburpsuite() error {
	cmd := exec.Command("sudo", "apt", "install", "burpsuite", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkburpsuiteDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
