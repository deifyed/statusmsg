package main

import (
	"bytes"
	"fmt"
	"github.com/deifyed/statusmsg/pkg/battery"
	"github.com/deifyed/statusmsg/pkg/clock"
	"github.com/deifyed/statusmsg/pkg/gme"
	"github.com/deifyed/statusmsg/pkg/update"
	"github.com/deifyed/statusmsg/pkg/volume"
	"log"
	"os"
	"path"
)

func main() {
	logPath := path.Join("/tmp", "statusbar.log")
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(fmt.Sprintf("unable to open logfile %s", logPath), err)
	}

	defer func() {
		_ = logFile.Close()
	}()

	log.SetOutput(logFile)

	volumeStatus, err := volume.GetStatus()
	if err != nil {
		log.Println(fmt.Errorf("error getting volume status: %w", err))
	}

	batteryStatus, err := battery.GetStatus()
	if err != nil {
		log.Println(fmt.Errorf("error getting battery status: %w", err))
	}

	updateStatus, _ := update.GetStatus()

	clockStatus := clock.GetStatus()

	gmeStatus, _ := gme.GetStatus()

	var buf bytes.Buffer

	if gmeStatus.Timestamp != 0 {
		fmt.Fprintf(&buf, "%s ", gmeStatus.String())
	}

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
