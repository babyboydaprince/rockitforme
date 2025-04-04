package installers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	_ "rockitforme/utils"
)

const burpsuiteCommand = "BurpSuiteCommunity"

func BurpsuiteInstall(check string, OpSystem string) bool {
	switch check {
	case "installed":
		if isburpsuiteInstalled() {
			return true
		} else {
			if err := installburpsuite(OpSystem); err != nil {
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

func isburpsuiteInstalled() bool {
	_, err := exec.LookPath(burpsuiteCommand)
	return err == nil
}

func installburpsuite(OpSystem string) error {
	switch OpSystem {
	case "debian":
		getBurp := exec.Command("wget", "-O",
			"/home/$USER/Downloads/Burpsuite_2025_1_5.sh",
			"https://portswigger-cdn.net/burp/releases/download?product=community&version=2025.1.5&type=Linux")
		getBurp.Stdout = os.Stdout
		getBurp.Stderr = os.Stderr
		err := getBurp.Run()
		if err != nil {
			return err
		}

		setExecPermission := exec.Command("chmod", "+x",
			"/home/$USER/Downloads/Burpsuite_2025_1_5.sh")
		setExecPermission.Stdout = os.Stdout
		setExecPermission.Stderr = os.Stderr
		permitErr := setExecPermission.Run()
		if permitErr != nil {
			return permitErr
		}

		setUpBurp := exec.Command("sudo", "bash",
			"/home/$USER/Downloads/Burpsuite_2025_1_5.sh")
		setUpBurp.Stdout = os.Stdout
		setUpBurp.Stderr = os.Stderr
		setUpErr := setUpBurp.Run()
		if setUpErr != nil {
			return setUpErr
		}

		return setUpBurp.Run()
	case "fedora":
		getBurp := exec.Command("wget", "-O",
			"/home/$USER/Downloads/Burpsuite_2025_1_5.sh",
			"https://portswigger-cdn.net/burp/releases/download?product=community&version=2025.1.5&type=Linux")
		getBurp.Stdout = os.Stdout
		getBurp.Stderr = os.Stderr
		err := getBurp.Run()
		if err != nil {
			return err
		}

		setExecPermission := exec.Command("chmod", "+x",
			"/home/$USER/Downloads/Burpsuite_2025_1_5.sh")
		setExecPermission.Stdout = os.Stdout
		setExecPermission.Stderr = os.Stderr
		permitErr := setExecPermission.Run()
		if permitErr != nil {
			return permitErr
		}

		setUpBurp := exec.Command("sudo", "bash",
			"/home/$USER/Downloads/Burpsuite_2025_1_5.sh")
		setUpBurp.Stdout = os.Stdout
		setUpBurp.Stderr = os.Stderr
		setUpErr := setUpBurp.Run()
		if setUpErr != nil {
			return setUpErr
		}

		return setUpBurp.Run()
	case "arch":
		projectRoot, err := getProjectMainDir()
		if err != nil {
			return fmt.Errorf("error getting project root: %w", err)
		}

		burpModulesPath := filepath.Join(projectRoot, "initialize", "installers", "modules", "burpsuite")

		if _, err := os.Stat(burpModulesPath); os.IsNotExist(err) {
			preparePath := exec.Command("mkdir", "-p",
				burpModulesPath)
			preparePath.Stdout = os.Stdout
			preparePath.Stderr = os.Stderr
			err = preparePath.Run()
			if err != nil {
				return fmt.Errorf("error making burpsuite module directory: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("error checking burpsuite module directory: %w", err)
		}

		getBurp := exec.Command("wget", "-O",
			filepath.Join(burpModulesPath, "Burpsuite_2025_1_5.sh"),
			"https://portswigger-cdn.net/burp/releases/download?product=community&version=2025.1.5&type=Linux")
		getBurp.Stdout = os.Stdout
		getBurp.Stderr = os.Stderr
		err = getBurp.Run()
		if err != nil {
			return err
		}

		setExecPermission := exec.Command("chmod", "+x",
			"/home/$USER/Downloads/Burpsuite_2025_1_5.sh")
		setExecPermission.Stdout = os.Stdout
		setExecPermission.Stderr = os.Stderr
		permitErr := setExecPermission.Run()
		if permitErr != nil {
			return permitErr
		}

		setUpBurp := exec.Command("sudo", "bash",
			"/home/$USER/Downloads/Burpsuite_2025_1_5.sh")
		setUpBurp.Stdout = os.Stdout
		setUpBurp.Stderr = os.Stderr
		setUpErr := setUpBurp.Run()
		if setUpErr != nil {
			return setUpErr
		}

		return setUpBurp.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func getProjectMainDir() (string, error) {
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
