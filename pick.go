// +build !linux
// +build !windows

package sfp

import (
	"fmt"
)

func pickMultiple() ([]string, error) {
	return "", fmt.Errorf("Sorry, pickMultiple is not supported for your platform")
}

func pick() (string, error) {
	return "", fmt.Errorf("Sorry, pick is not supported for your platform")
}
