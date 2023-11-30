package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const nfctoolsCommand = "nfc-scan-device"

func NfcToolsInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checknfctoolsDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isnfctoolsInstalled() {
			//fmt.Println("nfctools is already installed.")

			return true

		} else {
			//fmt.Println("Installing nfctools...")
			if err := installnfctools(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("nfctools installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isnfctoolsInstalled() bool {
	_, err := exec.LookPath(nfctoolsCommand)
	return err == nil
}

func installnfctools() error {
	cmd := exec.Command("sudo", "apt", "install", "nfctools", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checknfctoolsDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
