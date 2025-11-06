// Package battery handles retrieving information about the battery
package battery

import (
	"os"
	"path"
	"strings"
)

const defaultBatteryPath = "/sys/class/power_supply/BAT0"

var (
	defaultBatteryCapacityPath = path.Join(defaultBatteryPath, "capacity")
	defaultBatteryStatusPath   = path.Join(defaultBatteryPath, "status")
)

func Percentage(log logger) string {
	rawCapacity, err := os.ReadFile(defaultBatteryCapacityPath)
	if err != nil {
		log.Warnf("reading capacity: %s", err.Error())

		return "err"
	}

	return strings.TrimSpace(string(rawCapacity))
}

func Status(log logger) string {
	// #nosec G304 -- Defined above
	rawStatus, err := os.ReadFile(defaultBatteryStatusPath)
	if err != nil {
		log.Warnf("reading status: %s", err.Error())

		return "err"
	}

	discharging := strings.Contains(strings.ToLower(string(rawStatus)), "discharging")

	if discharging {
		return "-"
	}

	return "+"
}
