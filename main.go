package main

import (
	"github.com/ArpitChinmay/interview/setup"
)

func main() {
	// The SetupServer() method opens a db connection
	// and registers the routes.
	setup.SetupServer()
}
