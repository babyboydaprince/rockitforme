package installers

import (
	"fmt"
	"os"
	"os/exec"
	_ "rockitforme/utils"
)

const niktoCommand = "nikto"

// TODO - EXAMPLE OF DEPENDENCY SET UP
var niktoDependencies = map[string][]string{
	"debian": {"sudo", "ruby"},
	"fedora": {"sudo", "ruby"},
	"arch":   {"sudo", "ruby"},
}

func NiktoInstall(check string, OpSystem string) bool {
	switch check {
	case "dependencies":
		deps, ok := niktoDependencies[OpSystem]
		if !ok {
			return false
		}
		if err := checkniktoDependencies(deps); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isniktoInstalled() {
			return true
		} else {
			if err := installnikto(OpSystem); err != nil {
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

func isniktoInstalled() bool {
	_, err := exec.LookPath(niktoCommand)
	return err == nil
}

func installnikto(OpSystem string) error {
	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "nikto", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		cmd := exec.Command("sudo", "dnf", "install", "nikto", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "nikto")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func checkniktoDependencies(dependencies []string) error {
	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}
	return nil
}
