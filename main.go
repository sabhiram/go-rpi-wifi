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

	w, err := wifi.New("wlan0")
	fatalOnErr(err)

	w.DoTest()

	/*

	   1. Detect if we have wifi -- if we do, done
	   2. If we don't -- setup AP
	   3. Start server and host static files

	*/
}

func init() {
	log.SetPrefix("")
	log.SetFlags(0)
}

////////////////////////////////////////////////////////////////////////////////
