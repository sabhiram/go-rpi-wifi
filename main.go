package main

////////////////////////////////////////////////////////////////////////////////

import (
	"log"

	"github.com/sabhiram/go-rpi-wifi/wifi"
)

////////////////////////////////////////////////////////////////////////////////

func fatalOnErr(err error) {
	if err != nil {
		log.Fatalf(`Fatal error encountered : %s.
Program Aborting
`, err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	fatalOnErr(checkDependencies())

	w, err := wifi.New("wlan0", "rpi-config-ap")
	fatalOnErr(err)

	fatalOnErr(w.RescanInfo())

	if w.IsConnectedToNetwork() {
		log.Printf("Connected to wireless network with IP: %s\n", w.GetIP())
	} else {
		log.Printf("Not connected to WIFI - TOOD: Enable AP here!\n")
	}
}

func init() {
	log.SetPrefix("")
	log.SetFlags(0)
}

////////////////////////////////////////////////////////////////////////////////
