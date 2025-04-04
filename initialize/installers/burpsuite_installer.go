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
		homeDir, _ := os.UserHomeDir()
		downloadPath := filepath.Join(homeDir, "Downloads", "Burpsuite_2025_1_5.sh")
		getBurp := exec.Command("wget", "-O",
			downloadPath,
			"https://portswigger-cdn.net/burp/releases/download?product=community&version=2025.1.5&type=Linux")
		getBurp.Stdout = os.Stdout
		getBurp.Stderr = os.Stderr
		err := getBurp.Run()
		if err != nil {
			return err
		}

		setExecPermission := exec.Command("chmod", "+x",
			downloadPath)
		setExecPermission.Stdout = os.Stdout
		setExecPermission.Stderr = os.Stderr
		permitErr := setExecPermission.Run()
		if permitErr != nil {
			return permitErr
		}

		setUpBurp := exec.Command("sudo", "bash",
			downloadPath)
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
