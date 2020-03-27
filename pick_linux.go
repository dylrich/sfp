package sfp

import (
	"bufio"
	"io/ioutil"
	"os/exec"
	"strings"
)

func pickMultiple() ([]string, error) {
	var paths []string

	// TODO: support non-GTK environments
	cmd := exec.Command("zenity", "--file-selection", "--multiple", "--separator=\n")

	stdout, err := run(cmd)
	if err != nil {
		return paths, err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return paths, err
	}

	return paths, nil
}

func pick() (string, error) {
	// TODO: support non-GTK environments
	cmd := exec.Command("zenity", "--file-selection")

	stdout, err := run(cmd)
	if err != nil {
		return "", err
	}

	path, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	// zenity returns with a trailing \n by default
	return strings.TrimSpace(string(path)), nil
}
