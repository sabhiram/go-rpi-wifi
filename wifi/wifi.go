package wifi

////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"regexp"

	"github.com/sabhiram/go-rpi-wifi/exec"
)

////////////////////////////////////////////////////////////////////////////////

// Wifi represents this device's wifi information.
type Wifi struct {
	// Inputs
	iface string // interface name that we wish to operate on

	// Calculated fields
	HWAddr         string
	InetAddr       string
	APAddr         string
	ESSID          string
	IsUnassociated bool
}

func New(iface string) (*Wifi, error) {
	return &Wifi{
		iface: iface,
	}, nil
}

////////////////////////////////////////////////////////////////////////////////
/*
 *  Helper / internal functions.
 */

// updateWifiInfo runs `ifconfig` and `iwconfig` on the underlying interface
// for `w` and populates various fields for it.
func (w *Wifi) updateWifiInfo() error {
	var (
		err    error
		stdout []byte
	)

	// Run ifconfig
	stdout, _, err = exec.RunCommand("ifconfig", w.iface)
	if err != nil {
		return err
	} else {
		// Extract the hw address
		reHwAddr := regexp.MustCompile(`HWaddr\s([^\s]+)`)
		if m := reHwAddr.FindAllSubmatch(stdout, -1); m != nil {
			w.HWAddr = string(m[0][1])
		}

		// Extract the inet address
		reInetAddr := regexp.MustCompile(`inet addr:([^\s]+)`)
		if m := reInetAddr.FindAllSubmatch(stdout, -1); m != nil {
			w.InetAddr = string(m[0][1])
		}
	}

	// Run iwconfig
	stdout, _, err = exec.RunCommand("iwconfig", w.iface)
	if err != nil {
		return err
	} else {
		// Extract the access point name (upstream)
		reAPAddr := regexp.MustCompile(`Access Point:\s([^\s]+)`)
		if m := reAPAddr.FindAllSubmatch(stdout, -1); m != nil {
			w.APAddr = string(m[0][1])
		}

		// Extract the ESSID
		reESSID := regexp.MustCompile(`ESSID:\"([^\"]+)\"`)
		if m := reESSID.FindAllSubmatch(stdout, -1); m != nil {
			w.ESSID = string(m[0][1])
		}

		// Check is un-associated still
		reUnassoc := regexp.MustCompile(`(unassociated)\s+Nick`)
		if m := reUnassoc.FindAllSubmatch(stdout, -1); m != nil {
			w.IsUnassociated = true
		}
	}

	return nil
}
func (w *Wifi) restartWifi() error { return nil }
func (w *Wifi) scan() error        { return nil }

////////////////////////////////////////////////////////////////////////////////
/*
 *  Public APIs for the wifi package.
 */

// Return true if we have an external wireless connection, false otherwise.
func (w *Wifi) IsConnectedToNetwork() (bool, error) { return false, nil }

// Return true if we are currently running as an access point
func (w *Wifi) IsAccessPoint() (bool, error) { return false, nil }

////////////////////////////////////////////////////////////////////////////////

func (w *Wifi) DoTest() {
	fmt.Printf("Before update: %#v\n", w)
	w.updateWifiInfo()
	fmt.Printf("After update:  %#v\n", w)
}

////////////////////////////////////////////////////////////////////////////////
