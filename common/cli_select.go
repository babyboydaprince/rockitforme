// common/common.go

package common

import (
	"bufio"
	"fmt"
	"os"
	"rockitforme/banner"
	"strconv"
)

// SingleSelect displays a menu and prompts the user to select an option.
// It returns the selected option or an empty string if the user presses Enter without choosing.
func SingleSelect(prompt string, options []string) string {

	fmt.Println(prompt)

	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nChoose the tool set to work with: ")

	scanner.Scan()

	input := scanner.Text()

	if input == "" {

		fmt.Print("\033[H\033[2J") // Clear the console
		banner.PrintBanner()

		return ""
	}

	// Convert the input to an integer index
	choiceIndex, err := strconv.Atoi(input)
	if err != nil || choiceIndex < 1 || choiceIndex > len(options) {

		// Handle invalid input (not a number or out of range)
		fmt.Print("\033[H\033[2J") // Clear the console
		banner.PrintBanner()

		return ""
	}

	// Adjust the index since arrays are zero-based
	choiceIndex--

	return options[choiceIndex]
}
