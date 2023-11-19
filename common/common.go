package common

import (
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
