package initialize

import (
	"fmt"
	"rockitforme/banner"
	"rockitforme/initialize/installers"
	"time"

	"github.com/fatih/color"
)

func CheckForDependencies(OpSystem string) {
	fmt.Print("\033[H\033[2J")
	banner.PrintBanner()

	color.Yellow("\n\nChecking installed apps:\n\n")
	fmt.Println()

	checkAppsInstalled(OpSystem)

	time.Sleep(1 * time.Second)

	color.Yellow("\nStarting dependency check:\n\n")
	fmt.Println()

	// checkDependencies()

}

func checkDependencies() {
	// s := spinner.New(spinner.CharSets[15], 50*time.Millisecond)
	// s.Start()
	time.Sleep(3 * time.Second)
	// s.Stop()

	fmt.Print("Dependencies installed.....")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)
}

func checkAppsInstalled(OpSystem string) {
	// s := spinner.New(spinner.CharSets[15], 50*time.Millisecond)

	// s.Start()
	time.Sleep(500 * time.Millisecond)
	installers.NmapInstall("installed", OpSystem)
	fmt.Printf("\r%s NMAP.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.BettercapInstall("installed", OpSystem)
	fmt.Printf("\r%s Berttercap.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.AircrackInstall("installed", OpSystem)
	fmt.Printf("\r%s Aircrack-ng.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.DsniffInstall("installed", OpSystem)
	fmt.Printf("\r%s dsniff.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	//installers.OneshotInstall("installed")
	//
	//fmt.Printf("\r%sOneshot.....", "")
	//color.Green("OK")
	//fmt.Println()
	//time.Sleep(1 * time.Second)

	installers.BurpsuiteInstall("installed", OpSystem)
	fmt.Printf("\r%s Burpsuite.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.SqlmapInstall("installed", OpSystem)
	fmt.Printf("\r%s SQLMAP.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.NiktoInstall("installed", OpSystem)
	fmt.Printf("\r%s NIKTO.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.WpscanInstall("installed", OpSystem)
	fmt.Printf("\r%s WPSCAN.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.GqrxInstall("installed", OpSystem)
	fmt.Printf("\r%s GQRX.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.TsharkInstall("installed", OpSystem)
	fmt.Printf("\r%s iwconfig.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.ReaverInstall("installed", OpSystem)
	fmt.Printf("\r%s reaver.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.BullyInstall("installed", OpSystem)
	fmt.Printf("\r%s bully.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.CowpattyInstall("installed", OpSystem)
	fmt.Printf("\r%s callpatty.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.HashcatInstall("installed", OpSystem)
	fmt.Printf("\r%s hashcat.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	//installers.PyritInstall("installed", OpSystem)
	//fmt.Printf("\r%s pyrit.....", "")
	//color.Green("OK")
	//fmt.Println()
	//time.Sleep(500 * time.Millisecond)

	//installers.HcxdumptoolInstall("installed", OpSystem)
	//fmt.Printf("\r%s hcxdumptool.....", "")
	//color.Green("OK")
	//fmt.Println()
	//time.Sleep(500 * time.Millisecond)
	//
	//installers.HcxpcaptoolInstall("installed", OpSystem)
	//fmt.Printf("\r%s hcxpcaptool.....", "")
	//color.Green("OK")
	//fmt.Println()
	//time.Sleep(500 * time.Millisecond)

	installers.NfcToolsInstall("installed", OpSystem)
	fmt.Printf("\r%s NFC-TOOLS.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.MfocInstall("installed", OpSystem)
	fmt.Printf("\r%s MFOC.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.MfcukInstall("installed", OpSystem)
	fmt.Printf("\r%s MFCUK.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.IwconfigInstall("installed", OpSystem)
	fmt.Printf("\r%s iwconfig.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// TODO - DEPENDENCY CHECK needs WIFITE setup
	//installers.WifiteInstall("installed", OpSystem)
	//fmt.Printf("\r%s Wifite.....", "")
	//color.Green("OK")
	//fmt.Println()
	//time.Sleep(500 * time.Millisecond)

	// TODO - DEPENDENCY CHECK needs WIFITE setup
	installers.AirgeddonInstall("installed", OpSystem)
	// s.Stop()
	fmt.Printf("\r%s Airgeddon.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

}
