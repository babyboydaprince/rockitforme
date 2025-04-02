package installers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	_ "rockitforme/utils"
)

const bully = "bully"

// TODO - EXAMPLE OF DEPENDENCY SET UP
var bullyDependencies = map[string][]string{
	"debian": {"gcc", "make", "libpcap-dev", "libtool"},
	"fedora": {"gcc-c++", "make", "libpcap-devel", "libtool"},
	"arch":   {"gcc", "make", "libpcap", "libtool"},
}

func BullyInstall(check string, OpSystem string) bool {
	switch check {
	case "installed":
		if isBullyInstalled() {
			return true
		} else {
			if err := checkBullyDependencies(bullyDependencies[OpSystem], OpSystem); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if err := installBully(OpSystem); err != nil {
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

func isBullyInstalled() bool {
	_, err := exec.LookPath(bully)
	return err == nil
}

func installBully(OpSystem string) error {
	projectRoot, err := getProjectRootPath()
	if err != nil {
		return fmt.Errorf("error getting project root: %w", err)
	}

	bullyModulesPath := filepath.Join(projectRoot, "initialize", "installers", "modules", "bully")
	bullyBuildPath := filepath.Join(projectRoot, "initialize", "installers", "modules", "bully", "src")

	switch OpSystem {
	case "debian":
		setupBully := exec.Command("sudo", "apt", "install", "bully", "-y")
		setupBully.Stdout = os.Stdout
		setupBully.Stderr = os.Stderr
		err := setupBully.Run()
		if err != nil {
			return fmt.Errorf("error installing bully: %w", err)
		}
		return nil
	case "fedora":
		if _, err := os.Stat(bullyModulesPath); os.IsNotExist(err) {
			getRepo := exec.Command("git", "clone",
				"https://github.com/kimocoder/bully.git", bullyModulesPath)
			getRepo.Stdout = os.Stdout
			getRepo.Stderr = os.Stderr
			err := getRepo.Run()
			if err != nil {
				return fmt.Errorf("error cloning bully repository: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("error checking bully directory: %w", err)
		}
		//
		//autoreconfigBully := exec.Command("autoreconf", "-i")
		//autoreconfigBully.Dir = bullyModulesPath
		//autoreconfigBully.Stdout = os.Stdout
		//autoreconfigBully.Stderr = os.Stderr
		//err := autoreconfigBully.Run()
		//if err != nil {
		//	return fmt.Errorf("error on autoreconf on bully module path: %w", err)
		//}
		//
		//runConfFile := exec.Command("./", "configure")
		//runConfFile.Dir = bullyModulesPath
		//runConfFile.Stdout = os.Stdout
		//runConfFile.Stderr = os.Stderr
		//err = runConfFile.Run()
		//if err != nil {
		//	return fmt.Errorf("error running configure file on bully module path: %w", err)
		//}

		buildTool := exec.Command("make")
		buildTool.Dir = bullyBuildPath
		buildTool.Stdout = os.Stdout
		buildTool.Stderr = os.Stderr
		err = buildTool.Run()
		if err != nil {
			return fmt.Errorf("error building bully: %w", err)
		}

		bullyMakeInstall := exec.Command("make", "install")
		bullyMakeInstall.Dir = bullyBuildPath
		bullyMakeInstall.Stdout = os.Stdout
		bullyMakeInstall.Stderr = os.Stderr
		err = bullyMakeInstall.Run()
		if err != nil {
			return fmt.Errorf("error installing bully: %w", err)
		}

		return nil
	case "arch":
		if _, err := os.Stat(bullyModulesPath); os.IsNotExist(err) {
			getRepo := exec.Command("git", "clone",
				"https://github.com/kimocoder/bully.git", bullyModulesPath)
			getRepo.Stdout = os.Stdout
			getRepo.Stderr = os.Stderr
			err := getRepo.Run()
			if err != nil {
				return fmt.Errorf("error cloning bully repository: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("error checking bully directory: %w", err)
		}

		buildTool := exec.Command("make")
		buildTool.Dir = bullyModulesPath
		buildTool.Stdout = os.Stdout
		buildTool.Stderr = os.Stderr
		err := buildTool.Run()
		if err != nil {
			return fmt.Errorf("error building bully: %w", err)
		}

		bullyMakeInstall := exec.Command("make", "install")
		bullyMakeInstall.Dir = bullyModulesPath
		bullyMakeInstall.Stdout = os.Stdout
		bullyMakeInstall.Stderr = os.Stderr
		err = bullyMakeInstall.Run()
		if err != nil {
			return fmt.Errorf("error building bully: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}

func getProjectRootPath() (string, error) {
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

func checkBullyDependencies(dependencies []string, OpSystem string) error {
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
			return fmt.Errorf("error intsalling BULLY dependencies: %w", err)
		}
		return nil
	case "fedora":
		fixMissingDependencies := exec.Command("sudo", append([]string{"dnf", "install", "-y"}, missingDeps...)...)
		fixMissingDependencies.Stdout = os.Stdout
		fixMissingDependencies.Stderr = os.Stderr
		err := fixMissingDependencies.Run()
		if err != nil {
			return fmt.Errorf("error intsalling BULLY dependencies: %w", err)
		}
		return nil
	case "arch":
		fixMissingDependencies := exec.Command("sudo", append([]string{"pacman", "-S", "--noconfirm"}, missingDeps...)...)
		fixMissingDependencies.Stdout = os.Stdout
		fixMissingDependencies.Stderr = os.Stderr
		err := fixMissingDependencies.Run()
		if err != nil {
			return fmt.Errorf("error intsalling BULLY dependencies: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", OpSystem)
	}
}
