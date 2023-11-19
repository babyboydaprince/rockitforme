package initialize

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func Checkboxes(label string, opts []string) []string {
	var res []string
	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	err := survey.AskOne(prompt, &res)
	if err != nil {
		return nil
	}

	return res
}

func StartMenu() {
	answers := Checkboxes(
		"Which are your favourite programming languages?",
		[]string{
			"C",
			"Python",
			"Java",
			"C++",
			"C#",
			"Visual Basic",
			"JavaScript",
			"PHP",
			"Assembly Language",
			"SQL",
			"Groovy",
			"Classic Visual Basic",
			"Fortran",
			"R",
			"Ruby",
			"Swift",
			"MATLAB",
			"Go",
			"Prolog",
			"Perl",
		},
	)
	s := strings.Join(answers, ", ")
	fmt.Println("Oh, I see! You like", s)
}
