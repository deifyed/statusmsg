package volume

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var (
	bluetoothActive          = regexp.MustCompile("Active: active")
	headphonesConnectedRegex = regexp.MustCompile("Connected: yes")
)

func (s Status) String() string {
	var icon string

	switch {
	case s.isMuted():
		return "🔇"
	case s.Device == "headphones":
		icon = "🎧"
	default:
		icon = "🔊"
	}

	return fmt.Sprintf("%s%s", icon, strings.TrimSpace(s.Level))
}

func isBluetoothActive() (bool, error) {
	output, err := exec.Command("systemctl", "status", "bluetooth").Output()
	if err != nil {
		return false, fmt.Errorf("fetching bluetoothd information: %w", err)
	}

	return bluetoothActive.Match(output), nil
}

func GetStatus() (status Status, err error) {
	var (
		pamixerOutput,
		bluetoothOutput []byte
	)

	pamixerOutput, err = exec.Command("pamixer", "--get-volume-human").Output()
	if err != nil {
		return Status{}, fmt.Errorf("fetching pamixer information: %w", err)
	}

	if active, _ := isBluetoothActive(); active {
		bluetoothOutput, err = exec.Command("bluetoothctl", "info", "CC:98:8B:94:9F:59").Output()
		if err != nil {
			return Status{}, fmt.Errorf("fetching bluetooth device information: %w", err)
		}

		if headphonesConnectedRegex.Match(bluetoothOutput) {
			status.Device = "headphones"
		}
	}

	status.Level = string(pamixerOutput)

	return status, nil
}
