package banner

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func PrintBanner() {
	asciiLogo1 := figure.NewColorFigure("Rock it", "elite", "red", true)
	asciiLogo2 := figure.NewColorFigure("For me !", "elite", "blue", true)

	fmt.Print("\033[H\033[2J")

	asciiLogo1.Print()
	asciiLogo2.Print()
}

func BannerMacChanger() {
	asciiLogo1 := figure.NewColorFigure("Mac", "elite", "green", true)
	asciiLogo2 := figure.NewColorFigure("Changer !", "elite", "green", true)

	fmt.Print("\033[H\033[2J")

	asciiLogo1.Print()
	asciiLogo2.Print()
}

func BannerNetTools() {
	asciiLogo1 := figure.NewColorFigure("Net", "elite", "yellow", true)
	asciiLogo2 := figure.NewColorFigure("Tools !", "elite", "yellow", true)

	fmt.Print("\033[H\033[2J")

	asciiLogo1.Print()
	asciiLogo2.Print()
}

func BannerPortScanner() {
	asciiLogo1 := figure.NewColorFigure("Port", "elite", "red", true)
	asciiLogo2 := figure.NewColorFigure("Scanner !", "elite", "red", true)

	fmt.Print("\033[H\033[2J")

	asciiLogo1.Print()
	asciiLogo2.Print()
}

func BannerTraceRoute() {
	asciiLogo1 := figure.NewColorFigure("Trace", "elite", "cyan", true)
	asciiLogo2 := figure.NewColorFigure("Route !", "elite", "cyan", true)

	fmt.Print("\033[H\033[2J")

	asciiLogo1.Print()
	asciiLogo2.Print()
}

func BannerArpSpoof() {
	asciiLogo1 := figure.NewColorFigure("Arp", "elite", "purple", true)
	asciiLogo2 := figure.NewColorFigure("Spoof !", "elite", "purple", true)

	fmt.Print("\033[H\033[2J")

	asciiLogo1.Print()
	asciiLogo2.Print()
}
