package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const sqlmapCommand = "sqlmap"

func SqlmapInstall(check string) bool {

	switch check {
	case "dependencies":
		if err := checksqlmapDependencies(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if issqlmapInstalled() {
			//fmt.Println("sqlmap is already installed.")

			return true

		} else {
			//fmt.Println("Installing sqlmap...")
			if err := installsqlmap(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			//fmt.Println("sqlmap installed successfully.")

			return false
		}
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func issqlmapInstalled() bool {
	_, err := exec.LookPath(sqlmapCommand)
	return err == nil
}

func installsqlmap() error {
	cmd := exec.Command("sudo", "apt", "install", "sqlmap", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checksqlmapDependencies() error {
	dependencies := []string{"sudo"}

	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}

	return nil
}
