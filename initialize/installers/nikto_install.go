package installers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	_ "rockitforme/utils"
)

func NiktoInstall(check string, OpSystem string) bool {
	switch check {
	case "installed":
		perlCommand := []string{"perl"}
		niktoCommand := []string{"nikto"}
		perlReturedCommand := isPerlInstalled(perlCommand)
		if OpSystem == "fedora" {
			projectRoot, err := getNiktoProjectRoot()
			if err != nil {
				fmt.Printf("error getting project root: %w", err)
				return false
			}

			niktoModulesPathExist := []string{filepath.Join(projectRoot, "initialize", "installers", "modules", "nikto")}
			niktoFedoraReturedCommand := isNiktoInstalled(niktoModulesPathExist, OpSystem)

			if !perlReturedCommand[0] {
				if err := installPerl(OpSystem); err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			}
			if !niktoFedoraReturedCommand[0] {
				if err := installNikto(OpSystem); err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}

			}
			if perlReturedCommand[0] && niktoFedoraReturedCommand[0] {
				return true
			}
			return false
		}
		niktoReturedCommand := isNiktoInstalled(niktoCommand, OpSystem)
		if !perlReturedCommand[0] {
			if err := installPerl(OpSystem); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		}
		if !niktoReturedCommand[0] {
			if err := installNikto(OpSystem); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

		}
		if perlReturedCommand[0] && niktoReturedCommand[0] {
				return true
		}
		return false
	default:
		fmt.Println("Invalid check argument. Please use 'dependencies' or 'installed'.")
		os.Exit(1)
	}
	return false
}

func isPerlInstalled(perlCmdList []string) []bool {
	var perlCmdReturn []bool

	for _, perlCommand := range perlCmdList {
		_, err := exec.LookPath(perlCommand)
		if err != nil {
			perlCmdReturn = append(perlCmdReturn, false)
		} else {
			perlCmdReturn = append(perlCmdReturn, true)
		}
	}
	return perlCmdReturn
}

func isNiktoInstalled(niktoCmdList []string, SysOp string) []bool {
	var niktoCmdReturn []bool
	if SysOp == "fedora" {
		_, err := os.Stat(niktoCmdList[0])
		if err == nil {
			niktoCmdReturn = append(niktoCmdReturn, true)
		}
		if os.IsNotExist(err) {
			niktoCmdReturn = append(niktoCmdReturn, false)
		}
		// Other errors (e.g. permission denied) → treat as not existing
		return niktoCmdReturn
	}

	for _, niktoCommand := range niktoCmdList {
		_, err := exec.LookPath(niktoCommand)
		if err != nil {
			niktoCmdReturn = append(niktoCmdReturn, false)
		} else {
			niktoCmdReturn = append(niktoCmdReturn, true)
		}
	}
	return niktoCmdReturn
}

func installPerl(OpSystem string) error {
	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "perl", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		// https://github.com/sullo/nikto.git
		cmd := exec.Command("sudo", "dnf", "install", "perl", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

		cmd2 := exec.Command("sudo", "dnf", "install", "perl-App-cpanminusb", "-y")
		cmd2.Stdout = os.Stdout
		cmd2.Stderr = os.Stderr
		cmd2.Run()

		cmd3 := exec.Command("sudo", "cpanm", "Net::SSLeay", "IO::Socket::SSL", "Getopt::Long", "JSON::PP", "MIME::Base64", "Digest::MD5", "Time::Local", "Time::HiRes", "POSIX", "Socket", "IO::Socket", "Timeout")
		cmd3.Stdout = os.Stdout
		cmd3.Stderr = os.Stderr
		return cmd3.Run()
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "perl")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func installNikto(OpSystem string) error {
	projectRoot, err := getNiktoProjectRoot()
	if err != nil {
		return fmt.Errorf("error getting project root: %w", err)
	}

	niktoModulesPath := filepath.Join(projectRoot, "initialize", "installers", "modules", "nikto")

	switch OpSystem {
	case "debian":
		cmd := exec.Command("sudo", "apt", "install", "nikto", "-y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "fedora":
		if _, err := os.Stat(niktoModulesPath); os.IsNotExist(err) {
			getRepo := exec.Command("git", "clone",
				"https://github.com/sullo/nikto.git", niktoModulesPath)
			getRepo.Stdout = os.Stdout
			getRepo.Stderr = os.Stderr
			err = getRepo.Run()
			if err != nil {
				return fmt.Errorf("error cloning nikto repository: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("error checking nikto directory: %w", err)
		}
	case "arch":
		cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", "nikto")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
	return err
}

func getNiktoProjectRoot() (string, error) {
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
