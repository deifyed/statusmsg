// Package sound exposes information regarding sound
package sound

import (
	"os/exec"
	"regexp"
	"strings"
)

var volumeRe = regexp.MustCompile(`\d+%`)

func Volume(log logger) string {
	volume, err := exec.Command("pactl",
		"get-sink-volume", "@DEFAULT_SINK@",
	).Output()
	if err != nil {
		log.Warnf("fetching pactl information: %s", err.Error())

		return "err"
	}

	cleanedVolume := volumeRe.Find(volume)

	return strings.ReplaceAll(string(cleanedVolume), "\n", "")
}

var relevantDeviceMACs = []string{
	"88:C9:E8:38:1D:CC",
	"CC:98:8B:94:9F:59",
}

func DeviceType(log logger) string {
	if active, _ := isBluetoothActive(); !active {
		return "speaker"
	}

	for _, mac := range relevantDeviceMACs {
		connected, err := isDeviceConnected(mac)
		if err != nil {
			log.Warnf("fetching bluetooth device information: %s", err.Error())

			return "err"
		}

		if connected {
			return "headphones"
		}
	}

	return "speaker"
}
