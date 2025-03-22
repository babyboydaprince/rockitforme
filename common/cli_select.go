package common

import (
	"bufio"
	"fmt"
	"os"
	"rockitforme/banner"
	"strconv"
)

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

		fmt.Print("\033[H\033[2J")
		banner.PrintBanner()

		return ""
	}

	choiceIndex, err := strconv.Atoi(input)
	if err != nil || choiceIndex < 1 || choiceIndex > len(options) {
		fmt.Print("\033[H\033[2J")
		banner.PrintBanner()

		return ""
	}

	choiceIndex--

	return options[choiceIndex]
}
