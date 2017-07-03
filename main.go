package main

////////////////////////////////////////////////////////////////////////////////

import (
	"log"
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	log.Printf("Hello go-rpi-wifi!")
	if err := checkDependencies(); err != nil {
		log.Fatalf("dependency error: %s\n", err.Error())
	}

	log.Printf("All deps good!\n")

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
