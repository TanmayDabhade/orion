package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Confirm prompts the user with a yes/no question.
func Confirm(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/N]: ", prompt)

	text, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	text = strings.ToLower(strings.TrimSpace(text))
	return text == "y" || text == "yes"
}
