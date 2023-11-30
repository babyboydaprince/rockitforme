package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const gqrxCommand = "gqrx"

func GqrxInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkgqrxDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isgqrxInstalled() {
			//fmt.Println("gqrx is already installed.")

			return true

		} else {
			//fmt.Println("Installing gqrx...")
			if err := installgqrx(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("gqrx installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isgqrxInstalled() bool {
	_, err := exec.LookPath(gqrxCommand)
	return err == nil
}

func installgqrx() error {
	cmd := exec.Command("sudo", "apt", "install", "gqrx-sdr", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkgqrxDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
