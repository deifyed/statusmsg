package sound

import (
	"fmt"
	"os/exec"
	"regexp"
)

var bluetoothActive = regexp.MustCompile("Active: active")
var headphonesConnectedRegex = regexp.MustCompile("Connected: yes")

func isBluetoothActive() (bool, error) {
	output, err := exec.Command("systemctl", "status", "bluetooth").Output()
	if err != nil {
		return false, fmt.Errorf("fetching bluetoothd information: %w", err)
	}

	return bluetoothActive.Match(output), nil
}

func isDeviceConnected(mac string) (bool, error) {
	output, err := exec.Command("bluetoothctl", "info", mac).Output()
	if err != nil {
		return false, fmt.Errorf("fetching bluetooth device information: %w", err)
	}

	return headphonesConnectedRegex.Match(output), nil
}
