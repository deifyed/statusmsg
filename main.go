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
	"github.com/deifyed/statusmsg/pkg/tickers"
	"github.com/sirupsen/logrus"
)

const logPath = "/tmp/statusmsg.log"

func main() {
	// #nosec G304 -- var defined above for readability
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		log.Println(fmt.Sprintf("unable to open logfile %s", logPath), err)
	}

	defer func() { _ = logFile.Close() }()

	log := logrus.New()
	configureLogger(log, logFile)

	var buf bytes.Buffer

	gmePercentage, err := tickers.GetCurrentPercentage("GME")
	if err != nil {
		log.Warn(fmt.Sprintf("Error getting GME percentage: %s", err.Error()))
		gmePercentage = "N/A"
	}

	line := []string{
		formatTicker("GME", gmePercentage),
		formatSound(sound.DeviceType(log), sound.Volume(log)),
		clock.DTG(),
	}

	batteryInfo := formatBattery(battery.Status(log), battery.Percentage(log))
	if batteryInfo != "ERR" {
		line = append(line, batteryInfo)
	}

	fmt.Fprint(&buf, strings.Join(line, " :: "))

	fmt.Print(buf.String())
}

func formatTicker(symbol string, percentage string) string {
	return fmt.Sprintf("%s(%s%%)", symbol, percentage)
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

func configureLogger(log *logrus.Logger, out io.Writer) {
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	log.SetOutput(out)
}
