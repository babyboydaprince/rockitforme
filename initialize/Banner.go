package initialize

import "github.com/common-nighthawk/go-figure"

func Banner() {

	asciiLogo1 := figure.NewColorFigure("Rock it", "Bloody", "red", true)
	asciiLogo2 := figure.NewColorFigure("Fotr me!", "Bloody", "blue", true)

	asciiLogo1.Print()
	asciiLogo2.Print()
}
