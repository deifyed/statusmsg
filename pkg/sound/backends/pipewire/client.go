package pipewire

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

func getPipeWireObjects() ([]pwDumpResponseObject, error) {
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

	var response []pwDumpResponseObject

	err = json.NewDecoder(&stdout).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}

	return response, nil
}
