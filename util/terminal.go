package util

import (
	"bufio"
	"fmt"
	"os"
)

// IsPiped returns true when data is piped into Stdin
func IsPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// AskBool shows a prompt to ask a closed question
func AskBool(prompt string) (bool, error) {
	fmt.Printf("> %s [y] ", prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return false, err
	}

	switch scanner.Text() {
	case "y", "yes", "ok", "1":
		return true, nil
	default:
		return false, nil
	}
}
