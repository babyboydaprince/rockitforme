package net_tools

import (
	"fmt"
	"os"
	"os/exec"
)

func Bettercap()  {

	cmd := exec.Command("gnome-terminal",
		"--", "bash", "-c", "nmap --help; exec bash")

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for command:", err)
	
}