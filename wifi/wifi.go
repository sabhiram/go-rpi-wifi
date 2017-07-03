package wifi

////////////////////////////////////////////////////////////////////////////////

import ()

////////////////////////////////////////////////////////////////////////////////

// Wifi represents this device's wifi information.
type Wifi struct {
	isAccessPoint
}

func New() (*Wifi, error) {
	return &Wifi{}, nil
}

////////////////////////////////////////////////////////////////////////////////
/*
 *  Helper / internal functions.
 */

func (w *Wifi) updateWifiInfo() error { return nil }
func (w *Wifi) restartWifi() error    { return nil }

////////////////////////////////////////////////////////////////////////////////
/*
 *  Public APIs for the wifi package.
 */

// Return true if we have an external wireless connection, false otherwise.
func (w *Wifi) IsConnectedToNetwork() (bool, error) { return false, nil }

// Return true if we are currently running as an access point
func (w *Wifi) IsAccessPoint() (bool, error) { return false, nil }

////////////////////////////////////////////////////////////////////////////////
