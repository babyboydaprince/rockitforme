package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func GetLinuxDistro() string {
	if _, err := exec.LookPath("apt-get"); err == nil {
		return "debian"
	}

	if _, err := exec.LookPath("dnf"); err == nil {
		return "fedora"
	}
	if _, err := exec.LookPath("yum"); err == nil {
		return "fedora"
	}

	if _, err := exec.LookPath("pacman"); err == nil {
		return "arch"
	}

	return "unknown"
}

func IsRootOrAdmin() bool {
	switch runtime.GOOS {
	case "windows":
		return isAdminWindows()
	case "linux":
		return isRootUnix()
	case "darwin":
		return isRootUnix()
	default:
		return false
	}
}

func isAdminWindows() bool {
	cmd := exec.Command("net", "session")
	err := cmd.Run()
	return err == nil
}

func isRootUnix() bool {
	return os.Geteuid() == 0
}
