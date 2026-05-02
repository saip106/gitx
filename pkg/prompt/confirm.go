package prompt

import (
	"fmt"
	"strings"
)

// Confirm asks the user for confirmation with a specific message.
func Confirm(message string, defaultValue bool) bool {
	choices := "y/N"
	if defaultValue {
		choices = "Y/n"
	}

	fmt.Printf("%s (%s): ", message, choices)

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return defaultValue
	}

	response = strings.ToLower(strings.TrimSpace(response))
	if response == "" {
		return defaultValue
	}

	if response == "y" || response == "yes" {
		return true
	}

	return false
}
