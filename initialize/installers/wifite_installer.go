package installers

import (
	"fmt"
	"os"
	"os/exec"
	_ "rockitforme/utils"
)

const wifiteCommand = "wifite"

func WifiteInstall(check string, OpSystem string) bool {
	switch check {
	case "dependencies":
		if err := checkWifiteDependencies(OpSystem); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isWifiteInstalled() {
			return true
		} else {
			if err := installWifite(OpSystem); err != nil {
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

func isWifiteInstalled() bool {
	_, err := exec.LookPath(wifiteCommand)
	return err == nil
}

func installWifite(OpSystem string) error {
	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "wifite", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		cmd := exec.Command("sudo", "dnf", "install", "wifite", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "wifite")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func checkWifiteDependencies(OpSystem string) error {
	tools := []string{"iwconfig", "ifconfig", "aircrack-ng", "tshark",
		"reaver", "bully", "cowpatty", "pyrit", "hashcat",
		"hcxdumptool", "hcxpcaptool"}

	dependencies := []string{"sudo", "python3"}
	dependencies = append(dependencies, tools...)

	for _, dep := range dependencies {
		if dep == "sudo" || dep == "python3" {
			_, err := exec.LookPath(dep)
			if err != nil {
				return fmt.Errorf("dependency not found: %s", dep)
			}
			continue
		}
		if !isToolInstalled(dep) {
			InstallTool(dep, OpSystem)
		}

	}

	return nil
}

func isToolInstalled(tool string) bool {
	_, err := exec.LookPath(tool)
	return err == nil
}

// InstallTool TODO - WIFITE dependency setup
func InstallTool(tool string, OpSystem string) {
	switch tool {
	case "iwconfig":
		TsharkInstall("installed", OpSystem)
	case "ifconfig":
		fmt.Println("ifconfig installation not yet implemented")
	case "aircrack-ng":
		AircrackInstall("installed", OpSystem)
	case "tshark":
		fmt.Println("tshark installation not yet implemented")
	case "reaver":
		fmt.Println("reaver installation not yet implemented")
	case "bully":
		fmt.Println("bully installation not yet implemented")
	case "cowpatty":
		fmt.Println("cowpatty installation not yet implemented")
	case "pyrit":
		fmt.Println("pyrit installation not yet implemented")
	case "hashcat":
		fmt.Println("hashcat installation not yet implemented")
	case "hcxdumptool":
		fmt.Println("hcxdumptool installation not yet implemented")
	case "hcxpcaptool":
		fmt.Println("hcxpcaptool installation not yet implemented")
	default:
		fmt.Printf("Tool '%s' installation not yet implemented\n", tool)
	}
}
