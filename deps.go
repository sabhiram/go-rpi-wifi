package main

////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

////////////////////////////////////////////////////////////////////////////////

// This file intends to verify and validate that all the external dependencies
// of this program are installed and appropriately setup.

////////////////////////////////////////////////////////////////////////////////

var (
	requiredBinaries = []string{"tn", "dhcpd", "hostapd", "iw"}
	requiredFiles    = []string{"/etc/init.d/isc-dhcp-server"}
)

////////////////////////////////////////////////////////////////////////////////

// runCmd executes a command (list of strings, ex: runCmd("ping", "-n", "4"))
// and returns the stdout, stderr and an execution error (if running the command
// goes badly).
func runCmd(cmd string, opts ...string) ([]byte, []byte, error) {
	// Setup the command, grab some pipes.
	c := exec.Command(cmd, opts...)
	errPipe, err := c.StderrPipe()
	if err != nil {
		return nil, nil, err
	}
	stdPipe, err := c.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	// Run it!
	if err := c.Start(); err != nil {
		return nil, nil, err
	}

	// Flush the pipes.
	stderr, err := ioutil.ReadAll(errPipe)
	if err != nil {
		return nil, nil, err
	}
	stdout, err := ioutil.ReadAll(stdPipe)
	if err != nil {
		return nil, nil, err
	}

	// Done!
	return stdout, stderr, err
}

////////////////////////////////////////////////////////////////////////////////

// checkDependencies returns an error if any dependency is missing.  It will
// fail on the very first missing dependency.
func checkDependencies() error {
	// Check to see that we have all required binaries installed.
	for _, bindep := range requiredBinaries {
		stdout, stderr, err := runCmd("which", bindep)
		if err != nil {
			return fmt.Errorf("unable to check dependency: %s (%s)", bindep, err.Error())
		}
		if len(stderr) > 0 {
			return fmt.Errorf("unknown error encountered: %s", err.Error())
		}
		if len(stdout) == 0 {
			return fmt.Errorf("binary dependency missing: %s", bindep)
		}
	}

	// Check to see if the expected files are found.
	for _, filedep := range requiredFiles {
		if _, err := os.Stat(filedep); os.IsNotExist(err) {
			return fmt.Errorf("file dependency missing: %s", filedep)
		}
	}

	// All good!
	return nil
}

////////////////////////////////////////////////////////////////////////////////
