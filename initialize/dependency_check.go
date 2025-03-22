package initialize

import (
	"fmt"
	"rockitforme/banner"
	"rockitforme/initialize/installers"
	"time"

	"github.com/briandowns/spinner"
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

	checkDependencies()

}

func checkDependencies() {
	s := spinner.New(spinner.CharSets[15], 50*time.Millisecond)
	s.Start()
	time.Sleep(3 * time.Second)
	s.Stop()

	fmt.Print("Dependencies installed.....")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)
}

func checkAppsInstalled(OpSystem string) {
	s := spinner.New(spinner.CharSets[15], 50*time.Millisecond)

	s.Start()
	installers.NmapInstall("installed", OpSystem)
	fmt.Printf("\r%sNMAP.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.BettercapInstall("installed", OpSystem)
	fmt.Printf("\r%sBerttercap.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.AircrackInstall("installed", OpSystem)
	fmt.Printf("\r%sAircrack-ng.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.DsniffInstall("installed", OpSystem)
	fmt.Printf("\r%sdsniff.....", "")
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
	fmt.Printf("\r%sBurpsuite.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.SqlmapInstall("installed", OpSystem)
	fmt.Printf("\r%sSQLMAP.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.NiktoInstall("installed", OpSystem)
	fmt.Printf("\r%sNIKTO.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.WpscanInstall("installed", OpSystem)
	fmt.Printf("\r%sWPSCAN.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.GqrxInstall("installed", OpSystem)
	fmt.Printf("\r%sGQRX.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.NfcToolsInstall("installed", OpSystem)
	fmt.Printf("\r%sNFC-TOOLS.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.MfocInstall("installed", OpSystem)
	fmt.Printf("\r%sMFOC.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.MfcukInstall("installed", OpSystem)
	fmt.Printf("\r%sMFCUK.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.IwconfigInstall("installed", OpSystem)
	fmt.Printf("\r%siwconfig.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.TsharkInstall("installed", OpSystem)
	fmt.Printf("\r%siwconfig.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.ReaverInstall("installed", OpSystem)
	fmt.Printf("\r%reaver.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.BullyInstall("installed", OpSystem)
	fmt.Printf("\r%bully.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.CowpattyInstall("installed", OpSystem)
	fmt.Printf("\r%callpatty.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.PyritInstall("installed", OpSystem)
	fmt.Printf("\r%pyrit.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.HashcatInstall("installed", OpSystem)
	fmt.Printf("\r%hashcat.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.HcxdumptoolInstall("installed", OpSystem)
	fmt.Printf("\r%hcxdumptool.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	installers.HcxpcaptoolInstall("installed", OpSystem)
	fmt.Printf("\r%hcxpcaptool.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// TODO - DEPENDENCY CHECK needs WIFITE setup
	installers.WifiteInstall("installed", OpSystem)
	fmt.Printf("\r%sWifite.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// TODO - DEPENDENCY CHECK needs WIFITE setup
	installers.AirgeddonInstall("installed", OpSystem)
	s.Stop()
	fmt.Printf("\r%sAirgeddon.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

}
