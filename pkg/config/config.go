package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateConfig() (err error) {
	if boolPrompt("Would you like to configure some services?") {

	} else {
		fmt.Println("No services will be configured")
		return
	}

	return
}

// stringPrompt asks for a string value using the label
func stringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label)
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

// boolPrompt asks for a string value using the label
func boolPrompt(label string) bool {
	r := bufio.NewReader(os.Stdin)
	// print question and options
	fmt.Fprint(os.Stderr, label+" (y/N):")
	input, _ := r.ReadString('\n')
	switch strings.TrimSpace(input) {
	case "y", "Y", "Yes", "yes":
		return true
	default:
		return false
	}
}
