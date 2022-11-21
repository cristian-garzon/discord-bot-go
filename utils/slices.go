package utils

import "strings"

func Contains(commands []string, value string) bool {
	for _, v := range commands {
		if v == strings.ToLower(value) {
			return true
		}
	}

	return false
}
