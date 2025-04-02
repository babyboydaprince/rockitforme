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
	"debian": {"ruby", "ruby-dev", "libcurl4-openssl-dev", "libxslt1-dev", "sqlite3", "libsqlite3-dev", "build-essential"},
	"fedora": {"ruby", "ruby-devel", "libcurl-devel", "libxslt-devel", "sqlite", "sqlite-devel", "gcc", "make"},
	"arch":   {"ruby", "curl", "libxslt", "sqlite", "base-devel"},
}

func WpscanInstall(check string, OpSystem string) bool {
	switch check {
	case "installed":
		if iswpscanInstalled() {
			return true
		} else {
			if err := checkwpscanDependencies(wpscanDependencies[OpSystem], OpSystem); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
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
		setupWpscan := exec.Command("sudo", "gem", "install", "wpscan")
		setupWpscan.Stdout = os.Stdout
		setupWpscan.Stderr = os.Stderr
		err := setupWpscan.Run()
		if err != nil {
			return fmt.Errorf("error intsalling WPSCAN: %w", err)
		}
		return nil
	case "fedora":
		setupWpscan := exec.Command("sudo", "gem", "install", "wpscan")
		setupWpscan.Stdout = os.Stdout
		setupWpscan.Stderr = os.Stderr
		err := setupWpscan.Run()
		if err != nil {
			return fmt.Errorf("error intsalling WPSCAN: %w", err)
		}
		return nil
	case "arch":
		setupWpscan := exec.Command("sudo", "gem", "install", "wpscan")
		setupWpscan.Stdout = os.Stdout
		setupWpscan.Stderr = os.Stderr
		err := setupWpscan.Run()
		if err != nil {
			return fmt.Errorf("error intsalling WPSCAN: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func checkwpscanDependencies(dependencies []string, OpSystem string) error {
	var missingDeps []string

	for _, dep := range dependencies {
		if _, err := exec.LookPath(dep); err != nil {
			missingDeps = append(missingDeps, dep)
		}
	}

	if len(missingDeps) == 0 {
		fmt.Println("All dependencies are installed.")
		return nil
	}

	fmt.Printf("Missing dependencies: %v\n", missingDeps)

	switch OpSystem {
	case "debian":
		fixMissingDependencies := exec.Command("sudo", append([]string{"apt", "install", "-y"}, missingDeps...)...)
		fixMissingDependencies.Stdout = os.Stdout
		fixMissingDependencies.Stderr = os.Stderr
		err := fixMissingDependencies.Run()
		if err != nil {
			return fmt.Errorf("error intsalling WPSCAN dependencies: %w", err)
		}
		return nil
	case "fedora":
		fixMissingDependencies := exec.Command("sudo", append([]string{"dnf", "install", "-y"}, missingDeps...)...)
		fixMissingDependencies.Stdout = os.Stdout
		fixMissingDependencies.Stderr = os.Stderr
		err := fixMissingDependencies.Run()
		if err != nil {
			return fmt.Errorf("error intsalling WPSCAN dependencies: %w", err)
		}
		return nil
	case "arch":
		fixMissingDependencies := exec.Command("sudo", append([]string{"pacman", "-S", "--noconfirm"}, missingDeps...)...)
		fixMissingDependencies.Stdout = os.Stdout
		fixMissingDependencies.Stderr = os.Stderr
		err := fixMissingDependencies.Run()
		if err != nil {
			return fmt.Errorf("error intsalling WPSCAN dependencies: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}
