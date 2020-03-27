// Package sfp provides utilities for cross-platform file selection using native dialogs
package sfp

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

// Pick shells out to the system's native file selection dialog and returns the selected filepath along with any errors encountered
func Pick() (string, error) {
	return pick()
}

// PickMultiple shells out to the system's native file selection dialog and returns all selected filepaths along with any errors encountered
func PickMultiple() ([]string, error) {
	return pickMultiple()
}

func run(cmd *exec.Cmd) (io.Reader, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return stdout, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return stdout, err
	}

	if err := cmd.Start(); err != nil {
		return stdout, err
	}

	e, err := ioutil.ReadAll(stderr)
	if err != nil {
		return stdout, err
	}

	// stderr should not report anything. If it does, bail
	if string(e) != "" {
		return stdout, fmt.Errorf(string(e))
	}

	return stdout, nil
}
