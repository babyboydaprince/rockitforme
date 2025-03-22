package installers

import (
	"fmt"
	"os"
	"os/exec"
	_ "rockitforme/utils"
)

const wpscanCommand = "wpscan"

// TODO - EXAMPLE OF DEPENDENCY SET UP
var wpscanDependencies = map[string][]string{
	"debian": {"sudo", "ruby"},
	"fedora": {"sudo", "ruby"},
	"arch":   {"sudo", "ruby"},
}

func WpscanInstall(check string, OpSystem string) bool {
	switch check {
	case "dependencies":
		deps, ok := wpscanDependencies[OpSystem]
		if !ok {
			return false
		}
		if err := checkwpscanDependencies(deps); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if iswpscanInstalled() {
			return true
		} else {
			if err := installwpscan(OpSystem); err != nil {
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

func iswpscanInstalled() bool {
	_, err := exec.LookPath(wpscanCommand)
	return err == nil
}

func installwpscan(OpSystem string) error {
	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "wpscan", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		cmd := exec.Command("sudo", "dnf", "install", "wpscan", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "wpscan")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func checkwpscanDependencies(dependencies []string) error {
	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}
	return nil
}
