package initialize

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func Banner() {

	asciiLogo1 := figure.NewColorFigure("Rock it", "larry3d", "red", true)
	asciiLogo2 := figure.NewColorFigure("For me!", "larry3d", "blue", true)

	fmt.Print("\033[H\033[2J")

	asciiLogo1.Print()
	asciiLogo2.Print()

}
