package wifi

////////////////////////////////////////////////////////////////////////////////

import (
	"regexp"
	"strings"

	"github.com/sabhiram/go-rpi-wifi/exec"
)

////////////////////////////////////////////////////////////////////////////////

type info struct {
	HWAddr         string
	InetAddr       string
	APAddr         string
	ESSID          string
	IsUnassociated bool
}

// Wifi represents this device's wifi information.
type Wifi struct {
	// Inputs
	iface  string // interface name that we wish to operate on
	apName string // the ap that we advertise as
	winfo  *info
}

func New(iface string, apName string) (*Wifi, error) {
	return &Wifi{
		iface:  iface,
		apName: apName,
		winfo:  nil,
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

	w.winfo = &info{}

	// Run ifconfig
	stdout, _, err = exec.RunCommand("ifconfig", w.iface)
	if err != nil {
		return err
	} else {
		// Extract the hw address
		reHwAddr := regexp.MustCompile(`HWaddr\s([^\s]+)`)
		if m := reHwAddr.FindAllSubmatch(stdout, -1); m != nil {
			w.winfo.HWAddr = string(m[0][1])
		}

		// Extract the inet address
		reInetAddr := regexp.MustCompile(`inet addr:([^\s]+)`)
		if m := reInetAddr.FindAllSubmatch(stdout, -1); m != nil {
			w.winfo.InetAddr = string(m[0][1])
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
			w.winfo.APAddr = string(m[0][1])
		}

		// Extract the ESSID
		reESSID := regexp.MustCompile(`ESSID:\"([^\"]+)\"`)
		if m := reESSID.FindAllSubmatch(stdout, -1); m != nil {
			w.winfo.ESSID = string(m[0][1])
		}

		// Check is un-associated still
		reUnassoc := regexp.MustCompile(`(unassociated)\s+Nick`)
		if m := reUnassoc.FindAllSubmatch(stdout, -1); m != nil {
			w.winfo.IsUnassociated = true
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
func (w *Wifi) IsConnectedToNetwork() bool {
	return !w.IsAccessPoint() && len(w.winfo.InetAddr) > 0 && (w.winfo.IsUnassociated == false)
}

// Return true if we are currently running as an access point.
func (w *Wifi) IsAccessPoint() bool {
	return (strings.ToLower(w.winfo.HWAddr) == strings.ToLower(w.winfo.APAddr) &&
		w.winfo.ESSID == w.apName)
}

// Return the currently set inet addr.
func (w *Wifi) GetIP() string {
	return w.winfo.InetAddr
}

////////////////////////////////////////////////////////////////////////////////

func (w *Wifi) RescanInfo() error {
	return w.updateWifiInfo()
}

////////////////////////////////////////////////////////////////////////////////
