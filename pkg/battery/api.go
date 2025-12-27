// Package battery handles retrieving information about the battery
package battery

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

const defaultBatteryPath = "/sys/class/power_supply/BAT0"

var (
	defaultBatteryCapacityPath = path.Join(defaultBatteryPath, "capacity")
	defaultBatteryStatusPath   = path.Join(defaultBatteryPath, "status")
)

func Percentage() (string, error) {
	rawCapacity, err := os.ReadFile(defaultBatteryCapacityPath)
	if err != nil {
		return "", fmt.Errorf("reading capacity: %w", err)
	}

	return strings.TrimSpace(string(rawCapacity)), nil
}

var reChargingStatus = regexp.MustCompile(`(?i)discharging`)

func Charging() (bool, error) {
	// #nosec G304 -- Defined above
	rawStatus, err := os.ReadFile(defaultBatteryStatusPath)
	if err != nil {
		return false, fmt.Errorf("reading status: %w", err)
	}

	return !reChargingStatus.Match(rawStatus), nil
}
