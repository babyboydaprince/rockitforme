package initialize

import (
	"rockitforme/common"
	"rockitforme/tools"
	"strings"
)

func StartMenu() {
	answers := common.Checkboxes(
		"\nChoose tool set to use:",
		[]string{
			"Net tools",
			"Wifi tools",
		},
	)
	s := strings.Join(answers, ", ")

	if s == "Net Tools" {
		tools.GoChangeMyMac()
	}
}
