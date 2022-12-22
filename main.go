package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/deifyed/statusmsg/pkg/battery"
	"github.com/deifyed/statusmsg/pkg/clock"
	"github.com/deifyed/statusmsg/pkg/sound"
	"github.com/sirupsen/logrus"
)

const logPath = "/tmp/statusmsg.log"

func main() {
	// #nosec G304 -- var defined above for readability
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Println(fmt.Sprintf("unable to open logfile %s", logPath), err)
	}

	defer func() { _ = logFile.Close() }()

	log := logrus.New()
	configureLogger(log, logFile)

	var buf bytes.Buffer

	fmt.Fprint(&buf, strings.Join([]string{
		formatSound(sound.DeviceType(log), sound.Volume(log)),
		formatBattery(battery.Status(log), battery.Percentage(log)),
		formatClock(clock.DTG()),
	}, " :: "))

	fmt.Print(buf.String())
}

func formatBattery(status string, percentage string) string {
	if percentage == "err" {
		return "ERR"
	}

	return fmt.Sprintf("BAT %s%s%%", status, percentage)
}

func formatSound(deviceType string, volume string) string {
	if volume == "err" {
		return "ERR"
	}

	return fmt.Sprintf("%s/%s", deviceType, volume)
}

func formatClock(dtg string) string {
	return fmt.Sprintf("DTG %s", dtg)
}

func configureLogger(log *logrus.Logger, out io.Writer) {
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	log.SetOutput(out)
}
