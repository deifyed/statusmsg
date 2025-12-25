// Package pipewire retrieves information from PipeWire
package pipewire

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type client struct{}

func (c client) GetVolume() (int, error) {
	objects, err := getPipeWireObjects()
	if err != nil {
		return -1, fmt.Errorf("getting objects: %w", err)
	}

	defaultAudioSinkName, err := getDefaultAudioSinkName(objects)
	if err != nil {
		return -1, fmt.Errorf("getting default audio sink name: %w", err)
	}

	defaultAudioSinkNode, err := getNodeByName(objects, defaultAudioSinkName)
	if err != nil {
		return -1, fmt.Errorf("getting default audio sink node: %w", err)
	}

	volume, err := getVolumeProp(defaultAudioSinkNode.Info.Params.Props)
	if err != nil {
		return -1, fmt.Errorf("getting volume prop: %w", err)
	}

	return volume, nil
}

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
