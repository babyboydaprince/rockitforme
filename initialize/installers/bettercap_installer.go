package installers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	_ "rockitforme/utils"
)

const bettercapCommand = "bettercap"

var bettercapDependencies = map[string][]string{
	"debian": {"which", "libusb-1.0-0-dev"},
	"fedora": {"which", "libusb1-devel"},
	"arch":   {"which", "libusb"},
}

func BettercapInstall(check string, OpSystem string) bool {
	switch check {
	case "dependencies":
		deps, ok := bettercapDependencies[OpSystem]
		if !ok {
			return false
		}
		if err := checkbettercapDependencies(deps); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "installed":
		if isbettercapInstalled() {
			return true
		} else {
			if err := installbettercap(OpSystem); err != nil {
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

func isbettercapInstalled() bool {
	_, err := exec.LookPath(bettercapCommand)
	return err == nil
}

func installbettercap(OpSystem string) error {
	projectRoot, err := getProjectRoot()
	if err != nil {
		return fmt.Errorf("error getting project root: %w", err)
	}

	bettercapModulesPath := filepath.Join(projectRoot, "initialize", "installers", "modules", "bettercap")

	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "libusb-1.0-0-dev", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}

		cmd2 := exec.Command("sudo", "apt", "install", "bettercap", "-y")
		cmd2.Stdout = os.Stdout
		cmd2.Stderr = os.Stderr
		return cmd2.Run()
	case "fedora":
		installLibusb := exec.Command("sudo", "dnf", "install", "libusb1-devel", "-y")
		installLibusb.Stdout = os.Stdout
		installLibusb.Stderr = os.Stderr
		err := installLibusb.Run()
		if err != nil {
			return err
		}

		if _, err := os.Stat(bettercapModulesPath); os.IsNotExist(err) {
			getRepo := exec.Command("git", "clone",
				"https://github.com/bettercap/bettercap.git", bettercapModulesPath)
			getRepo.Stdout = os.Stdout
			getRepo.Stderr = os.Stderr
			err = getRepo.Run()
			if err != nil {
				return fmt.Errorf("error cloning bettercap repository: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("error checking bettercap directory: %w", err)
		}

		buildTool := exec.Command("make", "build")
		buildTool.Dir = bettercapModulesPath
		buildTool.Stdout = os.Stdout
		buildTool.Stderr = os.Stderr
		err = buildTool.Run()
		if err != nil {
			return fmt.Errorf("error building bettercap: %w", err)
		}

		installTool := exec.Command("sudo", "make", "install")
		installTool.Dir = bettercapModulesPath
		installTool.Stdout = os.Stdout
		installTool.Stderr = os.Stderr
		err = installTool.Run()
		if err != nil {
			return fmt.Errorf("error installing bettercap: %w", err)
		}

		return nil
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "libusb")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}

		cmd2 := exec.Command("sudo", "pacman", "-S", "--noconfirm", "bettercap")
		cmd2.Stdout = os.Stdout
		cmd2.Stderr = os.Stderr
		return cmd2.Run()
	default:
		return fmt.Errorf("Sad as molten popsicle.. :( \n"+
			"Your Linux Distro is not supported. But if you appreciate \n"+
			"ROck it for me! well enough, you may oppen an issue and, "+
			"request support for your Distro. :)\n %s", OpSystem)
	}

}

func checkbettercapDependencies(dependencies []string) error {
	for _, dep := range dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return fmt.Errorf("dependency not found: %s", dep)
		}
	}
	return nil
}

func getProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if filepath.Base(cwd) == "rockitforme" {
			return cwd, nil
		}

		parent := filepath.Dir(cwd)
		if parent == cwd {
			return "", fmt.Errorf("project root not found")
		}
		cwd = parent
	}
}
