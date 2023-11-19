package initialize

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/nexidian/gocliselect"
)

func Banner() {
	asciiLogo1 := figure.NewColorFigure("Go change", "Bloody", "red", true)
	asciiLogo2 := figure.NewColorFigure("my mac!", "Bloody", "blue", true)

	asciiLogo1.Print()
	asciiLogo2.Print()
}

func Main() {

	menu := gocliselect.NewMenu("Chose a colour")

	menu.AddItem("Red", "red")
	menu.AddItem("Blue", "blue")
	menu.AddItem("Green", "green")
	menu.AddItem("Yellow", "yellow")
	menu.AddItem("Cyan", "cyan")

	choice := menu.Display()

	fmt.Printf("Choice: %s\n", choice)

}
