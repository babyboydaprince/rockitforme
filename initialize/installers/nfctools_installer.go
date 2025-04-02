package installers

import (
	"fmt"
	"os"
	"os/exec"
	_ "rockitforme/utils"
)

const nfctoolsCommand = "nfc-list"

func NfcToolsInstall(check string, OpSystem string) bool {
	switch check {
	case "installed":
		if isnfctoolsInstalled() {
			return true
		} else {
			if err := installnfctools(OpSystem); err != nil {
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

func isnfctoolsInstalled() bool {
	_, err := exec.LookPath(nfctoolsCommand)
	return err == nil
}

func installnfctools(OpSystem string) error {
	//projectRoot, err := getProjectRootFolder()
	//if err != nil {
	//	return fmt.Errorf("error getting project root: %w", err)
	//}

	//nfctoolsdModulesPath := filepath.Join(projectRoot, "initialize", "installers",
	//	"modules", "libnfc")

	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "nfctools", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		cmd := exec.Command("sudo", "dnf", "install", "nfctools", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "nfctools")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

//
//func getProjectRootFolder() (string, error) {
//	cwd, err := os.Getwd()
//	if err != nil {
//		return "", err
//	}
//	for {
//		if filepath.Base(cwd) == "rockitforme" {
//			return cwd, nil
//		}
//		parent := filepath.Dir(cwd)
//		if parent == cwd {
//			return "", fmt.Errorf("project root not found")
//		}
//		cwd = parent
//	}
//}
