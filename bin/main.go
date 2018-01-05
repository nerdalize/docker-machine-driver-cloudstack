package main

import (
	"github.com/nerdalize/docker-machine-driver-cloudstack"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(cloudstack.NewDriver("", ""))
}
