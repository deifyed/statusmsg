package main

import (
	"bytes"
	"fmt"
	"github.com/deifyed/statusmsg/pkg/battery"
	"github.com/deifyed/statusmsg/pkg/clock"
	"github.com/deifyed/statusmsg/pkg/update"
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

	updateStatus, err := update.GetStatus()
	if err != nil {
		log.Fatal(err)
	}

	clockStatus := clock.GetStatus()

	var buf bytes.Buffer

	if updateStatus.PackageCount != 0 {
		fmt.Fprintf(&buf, "%s ", updateStatus.String())
	}

	fmt.Fprintf(
		&buf,
		"%s %s %s",
		volumeStatus.String(),
		batteryStatus.String(),
		clockStatus.String(),
	)

	fmt.Print(buf.String())
}
