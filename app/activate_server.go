package main

import (
	"github.com/activate_server/activate_server"
)

func main() {
	opts := activate_server.NewOptions()
	opts.Store()

	server := activate_server.NewActivateServer(opts)

	server.Run()
}
