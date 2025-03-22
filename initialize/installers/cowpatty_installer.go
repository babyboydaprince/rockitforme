package installers

import (
	"fmt"
	"os"
	"os/exec"
	_ "rockitforme/utils"
)

const cowpatty = "cowpatty"

// TODO - EXAMPLE OF DEPENDENCY SET UP
var cowpattyDependencies = map[string][]string{
	"debian": {"sudo", "ruby"},
	"fedora": {"sudo", "ruby"},
	"arch":   {"sudo", "ruby"},
}

func CowpattyInstall(check string, OpSystem string) bool {
	switch check {
	case "dependencies":
		deps, ok := cowpattyDependencies[OpSystem]
		if !ok {
			return false
		}
		if err := checkCowpattyDependencies(deps); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isCowpattyInstalled() {
			return true
		} else {
			if err := installCowpatty(OpSystem); err != nil {
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

func isCowpattyInstalled() bool {
	_, err := exec.LookPath(cowpatty)
	return err == nil
}

func installCowpatty(OpSystem string) error {
	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "cowpatty", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		cmd := exec.Command("sudo", "dnf", "install", "cowpatty", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "cowpatty")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func checkCowpattyDependencies(dependencies []string) error {
	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}
	return nil
}
