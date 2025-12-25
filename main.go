package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	bat "github.com/deifyed/statusmsg/pkg/battery"
	"github.com/deifyed/statusmsg/pkg/clock"
	"github.com/deifyed/statusmsg/pkg/sound/backends/pipewire"
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

	fmt.Print(status(log))
}

func status(log *logrus.Logger) string {
	line := make([]string, 0)

	if status, err := sound(); err == nil {
		line = append(line, status)
	} else {
		log.Warnf("Unable to get sound status: %s", err.Error())
	}

	if status, err := battery(); err == nil {
		line = append(line, status)
	} else {
		log.Warnf("Unable to get battery status: %s", err.Error())
	}

	line = append(line, clock.DTG())

	return strings.Join(line, " :: ")
}

func battery() (string, error) {
	charging, err := bat.Charging()
	if err != nil {
		return "", fmt.Errorf("acquiring status: %w", err)
	}

	batteryStatus := "-"
	if charging {
		batteryStatus = "+"
	}

	batteryPercentage, err := bat.Percentage()
	if err != nil {
		return "", fmt.Errorf("acquiring percentage: %w", err)
	}

	return fmt.Sprintf("BAT %s%s%%", batteryStatus, batteryPercentage), nil
}

func sound() (string, error) {
	soundClient := pipewire.Client{}

	deviceType, err := soundClient.GetDevice()
	if err != nil {
		return "", fmt.Errorf("acquiring device type: %w", err)
	}

	deviceVolume, err := soundClient.GetVolume()
	if err != nil {
		return "", fmt.Errorf("acquiring volume: %w", err)
	}

	return fmt.Sprintf("%s/%d", deviceType, deviceVolume), nil
}

func configureLogger(log *logrus.Logger, out io.Writer) {
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	log.SetOutput(out)
}
