package main

import (
	"flag"
	"github.com/elramir/config"
	"github.com/elramir/server"
)

func main() {

	mode := flag.String("m", "", "server mode, options: debug, release, test")
	port := flag.String("p", "", "server port")

	flag.Parse()

	config.SetDefault()

	if *mode != "" {
		config.Mode = *mode
	}
	if *port != "" {
		// config.Core.Port = *port
		config.Port = *port
	}

	server.RunHTTPServer()
	// }

	// func main() {

	// ptservice.SetVersion(Version)

	// if *version {
	// 	ptservice.PrintVersion()
	// 	return
	// }

	// ptservice.Init()
	// ptservice.RunHTTPServer()
}
