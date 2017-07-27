package util

import (
	"os"
)

// IsPiped returns true when data is piped into Stdin
func IsPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
