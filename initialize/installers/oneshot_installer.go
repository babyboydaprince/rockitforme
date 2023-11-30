package installers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const oneshotCommand = "oneshot.py"

func OneshotInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checkoneshotDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isoneshotInstalled() {
			//fmt.Println("oneshot is already installed.")

			return true

		} else {
			//fmt.Println("Installing oneshot...")
			if err := installoneshot(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("oneshot installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isoneshotInstalled() bool {
	_, err := exec.LookPath(oneshotCommand)
	return err == nil
}

func installoneshot() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := "../modules/oneshot.py"

	sourceFilePath := filepath.Join(currentDir, filePath)

	destinationDir := "/usr/local/bin/"

	// Get the base filename from the source file path
	_, fileName := filepath.Split(sourceFilePath)

	// Create the destination file path
	destinationFilePath := filepath.Join(destinationDir, fileName)

	cmd := exec.Command("cp", sourceFilePath, destinationFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkoneshotDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
