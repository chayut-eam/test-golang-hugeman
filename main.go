package main

import (
	"github.com/chayut-eam/test-golang-hugeman/config"
)

func main() {

	apiServer := config.Bootstrap()
	defer config.Teardown()

	apiServer.Start()
}
