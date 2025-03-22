package installers

import (
	"fmt"
	"os"
	"os/exec"
	_ "rockitforme/utils"
)

const reaver = "reaver"

// TODO - EXAMPLE OF DEPENDENCY SET UP
var reaverDependencies = map[string][]string{
	"debian": {"sudo", "ruby"},
	"fedora": {"sudo", "ruby"},
	"arch":   {"sudo", "ruby"},
}

func ReaverInstall(check string, OpSystem string) bool {
	switch check {
	case "dependencies":
		deps, ok := reaverDependencies[OpSystem]
		if !ok {
			return false
		}
		if err := checkReaverDependencies(deps); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isReaverInstalled() {
			return true
		} else {
			if err := installReaver(OpSystem); err != nil {
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

func isReaverInstalled() bool {
	_, err := exec.LookPath(reaver)
	return err == nil
}

func installReaver(OpSystem string) error {
	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "reaver", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	//case "fedora":
	//	cmd := exec.Command("sudo", "dnf", "install", "reaver", "-y")
	//	cmd.Stdout = os.Stdout
	//	cmd.Stderr = os.Stderr
	//	return cmd.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "reaver")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func checkReaverDependencies(dependencies []string) error {
	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}
	return nil
}
