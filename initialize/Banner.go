package initialize

import "github.com/common-nighthawk/go-figure"

func Banner() {

	asciiLogo1 := figure.NewColorFigure("Go change", "slant", "cyan", true)
	asciiLogo2 := figure.NewColorFigure("my mac!", "slant", "yellow", true)

	asciiLogo1.Print()
	asciiLogo2.Print()
}
