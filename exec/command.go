package exec

////////////////////////////////////////////////////////////////////////////////

import (
	"io/ioutil"
	"os/exec"
)

////////////////////////////////////////////////////////////////////////////////

// RunCommand executes a command (list of strings, ex: RunCommand("ping", "-n", "4"))
// and returns the stdout, stderr and an execution error (if running the command
// goes badly).
func RunCommand(cmd string, opts ...string) ([]byte, []byte, error) {
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
