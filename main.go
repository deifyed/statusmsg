package main

import (
	"fmt"
	"github.com/deifyed/statusmsg/pkg/battery"
	"github.com/deifyed/statusmsg/pkg/clock"
	"github.com/deifyed/statusmsg/pkg/volume"
	"log"
)

func main() {
	volumeStatus, err := volume.GetStatus()
	if err != nil {
		log.Fatal(err)
	}

	batteryStatus, err := battery.GetStatus()
	if err != nil {
		log.Fatal(err)
	}

	clockStatus := clock.GetStatus()

	fmt.Printf(
		"%s %s %s",
		volumeStatus.String(),
		batteryStatus.String(),
		clockStatus.String(),
	)
}
