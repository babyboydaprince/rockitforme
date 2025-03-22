package installers

import (
	"fmt"
	"os"
	"os/exec"
	_ "rockitforme/utils"
)

const bully = "bully"

// TODO - EXAMPLE OF DEPENDENCY SET UP
var bullyDependencies = map[string][]string{
	"debian": {"sudo", "ruby"},
	"fedora": {"sudo", "ruby"},
	"arch":   {"sudo", "ruby"},
}

func BullyInstall(check string, OpSystem string) bool {
	switch check {
	case "dependencies":
		deps, ok := bullyDependencies[OpSystem]
		if !ok {
			return false
		}
		if err := checkBullyDependencies(deps); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isBullyInstalled() {
			return true
		} else {
			if err := installBully(OpSystem); err != nil {
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

func isBullyInstalled() bool {
	_, err := exec.LookPath(bully)
	return err == nil
}

func installBully(OpSystem string) error {
	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "bully", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		cmd := exec.Command("sudo", "dnf", "install", "bully", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "bully")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func checkBullyDependencies(dependencies []string) error {
	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}
	return nil
}
