package temperature

import (
	"fmt"
	"os/exec"
	"regexp"
)

const defaultCPUSensor = "k10temp-pci-00c3"

var tctlTemperatureRe = regexp.MustCompile(`Tctl:\s+\+?([\d.]+)°C`)

func CPU() (string, error) {
	raw, err := exec.Command("sensors", defaultCPUSensor).Output()
	if err != nil {
		return "", fmt.Errorf("running command: %w", err)
	}

	matches := tctlTemperatureRe.FindStringSubmatch(string(raw))
	if len(matches) < 2 {
		return "", fmt.Errorf("no %s temperature found", "Tctl")
	}

	return matches[1], nil
}
