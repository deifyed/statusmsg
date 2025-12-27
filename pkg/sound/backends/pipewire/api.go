// Package pipewire retrieves information from PipeWire
package pipewire

import (
	"fmt"
	"math"
	"regexp"
)

const (
	DeviceTypeHeadphones = "headphones"
	DeviceTypeSpeaker    = "speaker"
)

var bluetoothRe = regexp.MustCompile(`(?i)bluetooth`)

type Client struct{}

func (c Client) GetVolume() (int, error) {
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

	volume, err := getVolume(defaultAudioSinkNode.ID)
	if err != nil {
		return -1, fmt.Errorf("acquiring volume: %w", err)
	}

	return int(math.Round(volume * 100)), nil
}

func (c Client) GetDevice() (string, error) {
	objects, err := getPipeWireObjects()
	if err != nil {
		return "", fmt.Errorf("getting objects: %w", err)
	}

	defaultAudioSinkName, err := getDefaultAudioSinkName(objects)
	if err != nil {
		return "", fmt.Errorf("getting default audio sink name: %w", err)
	}

	if bluetoothRe.Match([]byte(defaultAudioSinkName)) {
		return DeviceTypeHeadphones, nil
	}

	return DeviceTypeSpeaker, nil
}
