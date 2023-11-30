package initialize

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"rockitforme/banner"
	"rockitforme/initialize/installers"
	"time"
)

func CheckForDependencies() {
	fmt.Print("\033[H\033[2J") // Clear the console
	banner.PrintBanner()

	color.Yellow("\n\nChecking installed apps:\n\n")
	fmt.Println()

	checkAppsInstalled()

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

	fmt.Print("Installed.....")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)
}

func checkAppsInstalled() {
	s := spinner.New(spinner.CharSets[15], 50*time.Millisecond)

	s.Start()
	fmt.Printf("\r%sNMAP.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.AircrackInstall("installed")

	fmt.Printf("\r%sAircrack-ng.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.DsniffInstall("installed")

	fmt.Printf("\r%sdsniff.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.WifiteInstall("installed")

	fmt.Printf("\r%sWifite.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	//installers.OneshotInstall("installed")
	//
	//fmt.Printf("\r%sOneshot.....", "")
	//color.Green("OK")
	//fmt.Println()
	//time.Sleep(1 * time.Second)

	installers.AirgeddonInstall("installed")

	fmt.Printf("\r%sAirgeddon.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.BurpsuiteInstall("installed")

	fmt.Printf("\r%sBurpsuite.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.SqlmapInstall("installed")

	fmt.Printf("\r%sSQLMAP.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.NiktoInstall("installed")

	fmt.Printf("\r%sNIKTO.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.WpscanInstall("installed")

	fmt.Printf("\r%sWPSCAN.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.GqrxInstall("installed")

	fmt.Printf("\r%sGQRX.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.NfcToolsInstall("installed")

	fmt.Printf("\r%sNFC-TOOLS.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.MfocInstall("installed")

	fmt.Printf("\r%sMFOC.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	installers.MfcukInstall("installed")
	s.Stop()
	fmt.Printf("\r%sMFCUK.....", "")
	color.Green("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)
}
