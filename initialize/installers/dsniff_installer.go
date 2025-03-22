package installers

import (
	"fmt"
	"os"
	"os/exec"
	_ "rockitforme/utils"
)

const dsniffCommand = "dsniff"

// TODO - EXAMPLE OF DEPENDENCY SET UP
var dsniffDependencies = map[string][]string{
	"debian": {"sudo", "ruby"},
	"fedora": {"sudo", "ruby"},
	"arch":   {"sudo", "ruby"},
}

func DsniffInstall(check string, OpSystem string) bool {

	switch check {
	case "dependencies":
		deps, ok := dsniffDependencies[OpSystem]
		if !ok {
			return false
		}
		if err := checkdsniffDependencies(deps); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isdsniffInstalled() {
			return true
		} else {
			if err := installdsniff(OpSystem); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
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

func installdsniff(OpSystem string) error {

	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "dsniff", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		cmd := exec.Command("sudo", "dnf", "install", "dsniff", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "dsniff")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func checkdsniffDependencies(dependencies []string) error {
	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}
	return nil
}
