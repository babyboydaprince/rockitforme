package installers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	_ "rockitforme/utils"
)

const sqlmapCommand = "sqlmap"
const snapCommand = "snap"

func SqlmapInstall(check string, OpSystem string) bool {
	switch check {
	case "installed":
		if issqlmapInstalled() {
			return true
		} else {
			if err := isSnapStoreInstalled(); err != true {
				err := installSnapStore(OpSystem)
				if err != nil {
					return false
				}
			}
			if err := installSqlmap(OpSystem); err != nil {
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

func issqlmapInstalled() bool {
	_, err := exec.LookPath(sqlmapCommand)
	return err == nil
}

func isSnapStoreInstalled() bool {
	_, err := exec.LookPath(snapCommand)
	return err == nil
}

func installSqlmap(OpSystem string) error {
	switch OpSystem {
	case "debian":
		setupSqlmap := exec.Command("sudo", "apt", "install", "sqlmap", "-y")
		setupSqlmap.Stdout = os.Stdout
		setupSqlmap.Stderr = os.Stderr
		err := setupSqlmap.Run()
		if err != nil {
			return fmt.Errorf("error installing sqlmap: %w", err)
		}
		return nil
	case "fedora":
		setupSqlmap := exec.Command("sudo", "snap", "install", "sqlmap")
		setupSqlmap.Stdout = os.Stdout
		setupSqlmap.Stderr = os.Stderr
		err := setupSqlmap.Run()
		if err != nil {
			return fmt.Errorf("error installing sqlmap: %w", err)
		}

		return nil
	case "arch":
		setupSqlmap := exec.Command("sudo", "snap", "install", "sqlmap")
		setupSqlmap.Stdout = os.Stdout
		setupSqlmap.Stderr = os.Stderr
		err := setupSqlmap.Run()
		if err != nil {
			return fmt.Errorf("error installing sqlmap: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("Sad as molten popsicle.. :( \n"+
			"Your Linux Distro is not supported. But if you appreciate \n"+
			"ROck it for me! well enough, you may oppen an issue and, "+
			"request support for your Distro. :)\n %s", OpSystem)
	}

}

func installSnapStore(OpSystem string) error {
	projectRoot, err := getProjectRootDir()
	if err != nil {
		return fmt.Errorf("error getting project root: %w", err)
	}

	snapdModulesPath := filepath.Join(projectRoot, "initialize", "installers", "modules", "snapd")

	switch OpSystem {
	case "debian":
		setUpSnap := exec.Command("sudo", "apt", "install", "snapd", "-y")
		setUpSnap.Stdout = os.Stdout
		setUpSnap.Stderr = os.Stderr
		err = setUpSnap.Run()
		if err != nil {
			return fmt.Errorf("error cloning snapd repository: %w", err)
		}
		return nil
	case "fedora":
		setUpSnap := exec.Command("sudo", "dnf", "install", "snapd", "-y")
		setUpSnap.Stdout = os.Stdout
		setUpSnap.Stderr = os.Stderr
		err = setUpSnap.Run()
		if err != nil {
			return fmt.Errorf("error installing snapd: %w", err)
		}

		enableClassicSymLink := exec.Command("sudo", "ln", "-s",
			"/var/lib/snapd/snap", " /snap")
		enableClassicSymLink.Stdout = os.Stdout
		enableClassicSymLink.Stderr = os.Stderr
		err = enableClassicSymLink.Run()
		if err != nil {
			return fmt.Errorf("error installing snapd: %w", err)
		}
		return nil
	case "arch":
		if _, err := os.Stat(snapdModulesPath); os.IsNotExist(err) {
			getRepo := exec.Command("git", "clone",
				"https://aur.archlinux.org/snapd.git", snapdModulesPath)
			getRepo.Stdout = os.Stdout
			getRepo.Stderr = os.Stderr
			err = getRepo.Run()
			if err != nil {
				return fmt.Errorf("error cloning snapd repository: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("error checking snapd directory: %w", err)
		}

		buildTool := exec.Command("makepkg", "-si")
		buildTool.Dir = snapdModulesPath
		buildTool.Stdout = os.Stdout
		buildTool.Stderr = os.Stderr
		err = buildTool.Run()
		if err != nil {
			return fmt.Errorf("error building snapd: %w", err)
		}

		enableSnapd := exec.Command("sudo", "systemctl", "enable",
			"--now", "snapd.socket")
		enableSnapd.Stdout = os.Stdout
		enableSnapd.Stderr = os.Stderr
		err = enableSnapd.Run()
		if err != nil {
			return fmt.Errorf("error installing snapd: %w", err)
		}

		enableAppArmor := exec.Command("sudo", "systemctl", "enable",
			"--now", "snapd.apparmor.service")
		enableAppArmor.Stdout = os.Stdout
		enableAppArmor.Stderr = os.Stderr
		err = enableAppArmor.Run()
		if err != nil {
			return fmt.Errorf("error installing snapd: %w", err)
		}

		enableClassicSymLink := exec.Command("sudo", "ln", "-s",
			"/var/lib/snapd/snap", " /snap")
		enableClassicSymLink.Stdout = os.Stdout
		enableClassicSymLink.Stderr = os.Stderr
		err = enableClassicSymLink.Run()
		if err != nil {
			return fmt.Errorf("error installing snapd: %w", err)
		}

		setupSnapStoreWithSnapd := exec.Command("sudo", "snap", "install", "snap-store")
		setupSnapStoreWithSnapd.Stdout = os.Stdout
		setupSnapStoreWithSnapd.Stderr = os.Stderr
		err := setupSnapStoreWithSnapd.Run()
		if err != nil {
			return fmt.Errorf("error installing snapd: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func getProjectRootDir() (string, error) {
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
