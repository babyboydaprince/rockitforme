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
