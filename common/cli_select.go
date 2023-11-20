package common

import (
	"fmt"
)

func SingleSelect(label string, opts []string) string {
	fmt.Println(label)
	for i, opt := range opts {
		fmt.Printf("%d. %s\n", i+1, opt)
	}

	var choice int
	fmt.Print("\nEnter your choice: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		panic(err)
	}

	if choice < 1 || choice > len(opts) {
		fmt.Println("\nInvalid choice. Please enter a valid number.")
		return SingleSelect(label, opts)
	}

	return opts[choice-1]
}
