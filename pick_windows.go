package sfp

import (
	"bufio"
	"io/ioutil"
	"os/exec"
	"strings"
)

const (
	// it's windows, shrug
	psm = `Add-Type -AssemblyName System.Windows.Forms
$f = new-object Windows.Forms.OpenFileDialog
$f.InitialDirectory = pwd
$f.Multiselect = $true
[void]$f.ShowDialog()	
if ($f.Multiselect) { $f.FileNames } else { $f.FileName }`

	pss = `Add-Type -AssemblyName System.Windows.Forms
$f = new-object Windows.Forms.OpenFileDialog
$f.InitialDirectory = pwd
$f.Multiselect = $false
[void]$f.ShowDialog()	
if ($f.Multiselect) { $f.FileNames } else { $f.FileName }`
)

func pickMultiple() ([]string, error) {
	var paths []string

	// - is necessary to fix ps echo behavior
	cmd := exec.Command("powershell", "-")

	cmd.Stdin = strings.NewReader(psm)

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
	// - is necessary to fix ps echo behavior
	cmd := exec.Command("powershell", "-")

	cmd.Stdin = strings.NewReader(pss)

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

	return strings.ReplaceAll(string(path), "\r\n", ""), nil
}
