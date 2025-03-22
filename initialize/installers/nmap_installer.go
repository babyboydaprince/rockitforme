package installers

import (
	"fmt"
	"os"
	"os/exec"
)

const nmapCommand = "nmap"

// TODO - EXAMPLE OF DEPENDENCY SET UP
var nmapDependencies = map[string][]string{
	"debian": {"sudo", "ruby"},
	"fedora": {"sudo", "ruby"},
	"arch":   {"sudo", "ruby"},
}

func NmapInstall(check string, OpSystem string) bool {
	switch check {
	case "dependencies":
		deps, ok := nmapDependencies[OpSystem]
		if !ok {
			return false
		}
		if err := checkNmapDependencies(deps); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isNmapInstalled() {
			return true
		} else {
			if err := installNmap(OpSystem); err != nil {
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

func isNmapInstalled() bool {
	_, err := exec.LookPath(nmapCommand)
	return err == nil
}

func installNmap(OpSystem string) error {

	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "nmap", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		cmd := exec.Command("sudo", "dnf", "install", "nmap", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "nmap")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}

}

func checkNmapDependencies(dependencies []string) error {
	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}
	return nil
}
