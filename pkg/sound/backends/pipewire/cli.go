package pipewire

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

func sanitizeVolume(raw []byte) string {
	asString := string(raw)
	trimmed := strings.Trim(asString, "\n")

	return strings.Split(trimmed, " ")[1]
}

func getVolume(id int) (float64, error) {
	rawVolume, err := exec.Command("wpctl",
		"get-volume", fmt.Sprintf("%d", id),
	).Output()
	if err != nil {
		return -1, fmt.Errorf("running command: %w", err)
	}

	saneVolume := sanitizeVolume(rawVolume)

	volume, err := strconv.ParseFloat(saneVolume, 64)
	if err != nil {
		return -1, fmt.Errorf("parsing command output: %w", err)
	}

	return volume, nil
}

func getPipewireDump() (io.Reader, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd := exec.Command("pw-dump")

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("%w: %s", err, stderr.String())

		return nil, fmt.Errorf("running command: %w", err)
	}

	return &stdout, nil
}
