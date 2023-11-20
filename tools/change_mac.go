package tools

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"rockitforme/banner"
	"rockitforme/common"
	"runtime"
	"strings"
	"time"

	_ "github.com/AlecAivazis/survey/v2"
	_ "github.com/google/gopacket/pcap"
	_ "github.com/jedib0t/go-pretty/v6/table"
)

func getOriginalMAC(interfaceName string) (string, error) {

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("wmic", "nic", "where",
			fmt.Sprintf("NetConnectionID='%s'", interfaceName), "get", "MACAddress", "/format:list")

	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("cat", fmt.Sprintf("/sys/class/net/%s/address", interfaceName))

	} else {
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func restoreOriginalMAC(interfaceName, originalMAC string) error {
	return changeMAC(interfaceName, originalMAC)
}

func winRestoreOriginalMAC() {

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	if len(interfaces) > 0 {
		originalMAC := interfaces[0].HardwareAddr
		fmt.Printf("Original MAC address: %s\n", originalMAC)
	} else {
		log.Fatal("unable to retrieve original MAC address")
	}

	return
}

func setRandomMAC(interfaceName string) error {
	randMAC, err := generateRandomMAC()
	if err != nil {
		return err
	}

	return changeMAC(interfaceName, randMAC)
}

func generateRandomMAC() (string, error) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	firstByte := 0x02
	otherBytes := make([]byte, 5)
	r.Read(otherBytes)

	randomMAC := fmt.Sprintf(
		"%02x:%02x:%02x:%02x:%02x:%02x", firstByte,
		otherBytes[0], otherBytes[1], otherBytes[2], otherBytes[3], otherBytes[4])

	return randomMAC, nil
}

func winChangeMAC(interfaceName string, newMAC string) error {

	// Disable the network interface
	disableCmd := exec.Command("netsh", "interface", "set", "interface", interfaceName, "admin=disable")
	err := disableCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to disable network interface: %w", err)
	}

	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	// Make a packet to change the MAC address
	packetData := []byte{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Destination MAC (broadcast)
		newMAC[0], newMAC[1], newMAC[2], newMAC[3], newMAC[4], newMAC[5], // Source MAC
		0x08, 0x06, // EtherType: ARP
		0x00, 0x01, // Hardware Type: Ethernet (1)
		0x08, 0x00, // Protocol Type: IPv4 (0x0800)
		0x06,       // Hardware Address Length: 6
		0x04,       // Protocol Address Length: 4
		0x00, 0x02, // Operation: Reply (2)
		// Sender Hardware Address (Source MAC)
		newMAC[0], newMAC[1], newMAC[2], newMAC[3], newMAC[4], newMAC[5],
		// Sender Protocol Address (IPv4 address, e.g., 192.168.1.1)
		0x00, 0x00, 0x00, 0x00,
		// Target Hardware Address (Destination MAC)
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// Target Protocol Address (IPv4 address, e.g., 192.168.1.2)
		0x00, 0x00, 0x00, 0x00,
	}

	packet := gopacket.NewPacket(packetData, layers.LayerTypeEthernet, gopacket.Default)

	// Send the packet to change the MAC address
	err = handle.WritePacketData(packet.Data())
	if err != nil {
		return err
	}

	// Enable the network interface
	enableCmd := exec.Command("netsh", "interface", "set", "interface", interfaceName, "admin=enable")
	err = enableCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to enable network interface: %w", err)
	}

	fmt.Printf("Windows MAC address of %s changed to %s\n", interfaceName, newMAC)
	return nil

}

func changeMAC(interfaceName, newMAC string) error {
	var cmd *exec.Cmd

	if runtime.GOOS == "linux" {
		cmd = exec.Command("ip", "link", "set", interfaceName, "address", newMAC)

	} else {
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func findInterfaces() {
	banner.PrintBanner()

	t := table.NewWriter()

	t.SetTitle("Network interfaces")

	t.AppendHeader(table.Row{
		"#", "Name", "Index", "MTU", "MAC Address"})

	if runtime.GOOS == "linux" {
		interfaces, err := net.Interfaces()

		if err != nil {
			fmt.Println("Error: ", err)
			return

		}

		for i, iface := range interfaces {
			t.AppendRow(table.Row{
				i + 1, iface.Name, iface.Index, iface.MTU, iface.HardwareAddr})

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()
			fmt.Println(t.Render())
			time.Sleep(50 * time.Millisecond)
		}
		return

	} else if runtime.GOOS == "windows" {

		// Get the list of available interfaces
		devices, err := pcap.FindAllDevs()
		if err != nil {
			log.Fatal(err)
		}

		for i, device := range devices {
			t.AppendRow(table.Row{
				i + 1, device.Name, device.Description})

			fmt.Print("\033[H\033[2J") // Clear the console
			banner.PrintBanner()
			fmt.Println(t.Render())
			time.Sleep(500 * time.Millisecond)
		}

	}
}

func GoChangeMyMac() {

	fmt.Print("\033[H\033[2J") // Clear the console
	banner.PrintBanner()

	for {
		menuChoice := common.SingleSelect("\nChoose an operation:\n", []string{
			"List interfaces",
			"Restore original MAC for Windows",
			"Change MAC Windows compatible",
			"Restore original MAC",
			"Set random MAC",
			"Set MAC manually",
			"Exit",
		})

		switch menuChoice {
		case "List interfaces":
			findInterfaces()
		case "Restore original MAC for Windows":
			winRestoreOriginalMAC()
		case "Change MAC Windows compatible":
			interfaceName := askForInterface()
			newMAC := askForNewMAC()
			err := winChangeMAC(interfaceName, newMAC)
			handleError(err)
		case "Restore original MAC":
			interfaceToRestore := askForInterface()
			originalMAC, err := getOriginalMAC(interfaceToRestore)
			handleError(err)
			err = restoreOriginalMAC(interfaceToRestore, originalMAC)
			handleError(err)
			fmt.Printf("\nOriginal MAC address for %s restored: %s\n", interfaceToRestore, originalMAC)
		case "Set random MAC":
			interfaceName := askForInterface()
			err := setRandomMAC(interfaceName)
			handleError(err)
			fmt.Printf("\nRandomized MAC address set for %s\n", interfaceName)
		case "Set MAC manually":
			interfaceName := askForInterface()
			manualMAC := askForManualMAC()
			err := changeMAC(interfaceName, manualMAC)
			handleError(err)
			fmt.Printf("\nMAC address for %s changed to %s\n", interfaceName, manualMAC)
		case "Exit":
			fmt.Println("\nExiting...")
			os.Exit(0)
		}
	}
}

func askForInterface() string {
	var interfaceName string
	prompt := &survey.Input{
		Message: "\nEnter the interface name:",
	}
	err := survey.AskOne(prompt, &interfaceName, survey.WithValidator(survey.Required))
	handleError(err)
	return interfaceName
}

func askForNewMAC() string {
	var newMAC string
	prompt := &survey.Input{
		Message: "Enter the new MAC address (in the format XX:XX:XX:XX:XX:XX):",
	}
	err := survey.AskOne(prompt, &newMAC, survey.WithValidator(survey.Required))
	handleError(err)
	return newMAC
}

func askForManualMAC() string {
	var manualMAC string
	prompt := &survey.Input{
		Message: "Enter the new MAC address (in the format XX:XX:XX:XX:XX:XX):",
	}
	err := survey.AskOne(prompt, &manualMAC, survey.WithValidator(survey.Required))
	handleError(err)
	return manualMAC
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
