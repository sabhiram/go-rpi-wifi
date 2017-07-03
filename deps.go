package main

////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"strings"
	// "os"

	"github.com/sabhiram/go-rpi-wifi/exec"
)

////////////////////////////////////////////////////////////////////////////////

// This file intends to verify and validate that all the external dependencies
// of this program are installed and appropriately setup.

////////////////////////////////////////////////////////////////////////////////

var (
	requiredBinaries = []struct{ binName, packageName string }{
		{"dhcpcd", "dhcpd"},
		{"hostapd", "hostadp"},
		{"iw", "iw"},
	}
	requiredFiles = []string{"/etc/init.d/isc-dhcp-server"}
)

////////////////////////////////////////////////////////////////////////////////

// checkDependencies returns an error if any dependency is missing.  It will
// fail on the very first missing dependency.
func checkDependencies() error {
	// Check to see if the expected files are found.
	// for _, filedep := range requiredFiles {
	// 	if _, err := os.Stat(filedep); os.IsNotExist(err) {
	// 		return fmt.Errorf("file dependency missing: %s", filedep)
	// 	}
	// }

	// Check to see that we have all required binaries installed.
	missing := []string{}
	for _, bindep := range requiredBinaries {
		stdout, stderr, err := exec.RunCommand("which", bindep.binName)
		if err != nil {
			return fmt.Errorf("unable to check dependency: %s (%s)", bindep.binName, err.Error())
		}
		if len(stderr) > 0 {
			return fmt.Errorf("unknown error encountered: %s", err.Error())
		}
		if len(stdout) == 0 {
			missing = append(missing, bindep.packageName)
		}
	}

	if len(missing) > 0 {
		s := fmt.Sprintf("sudo apt-get install -y %s", strings.Join(missing, " "))
		return fmt.Errorf("missing dependencies, update using: \"%s\"", s)
	}

	// All good!
	return nil
}

////////////////////////////////////////////////////////////////////////////////
