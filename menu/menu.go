package menu

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/turret-io/go-menu/menu"
	_ "github.com/turret-io/go-menu/menu"
)

func banner() {

	asciiLogo1 := figure.NewColorFigure("Go change", "Bloody", "red", true)
	asciiLogo2 := figure.NewColorFigure("my mac!", "Bloody", "blue", true)

	asciiLogo1.Print()
	asciiLogo2.Print()
}

func option1(args ...string) error {
	// Do something
	fmt.Println("number 1")
	return nil
}

func option2(args ...string) error {
	// Do something
	fmt.Println("number 2")
	return nil
}

func option3(args ...string) error {
	// Do something
	fmt.Println("number 3")
	return nil
}

func option4(args ...string) error {
	// Do something
	fmt.Println("number 4")
	return nil
}

func option5(args ...string) error {
	// Do something
	fmt.Println("number 5")
	return nil
}

func handleOption(option string) {
	fmt.Printf("Selected option: %s\n", option)
	// Add your logic based on the selected option
}

func main() {
	menuOptions := []menu.CommandOption{
		{"Option 1", "First option description", option1},
		{"Option 2", "Second option description", option2},
		{"Option 3", "Third option description", option3},
		{"Option 4", "Fourth option description", option4},
		{"Option 5", "Fifth option description", option5},
	}

	startMenu := menu.NewMenu("Select an option", menuOptions)
	startMenu.Start()
}
