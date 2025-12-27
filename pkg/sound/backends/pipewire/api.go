// Package pipewire retrieves information from PipeWire
package pipewire

import (
	"fmt"
	"math"
	"regexp"

	"github.com/deifyed/statusmsg/pkg/sound/backends/pipewire/objects"
)

const (
	DeviceTypeHeadphones = "headphones"
	DeviceTypeSpeaker    = "speaker"
)

var bluetoothRe = regexp.MustCompile(`(?i)bluetooth`)

type Client struct{}

func (c Client) GetVolume() (int, error) {
	dump, err := getPipewireDump()
	if err != nil {
		return -1, fmt.Errorf("acquiring pipewire dump: %w", err)
	}

	defaultAudioSinkID, err := objects.GetDefaultAudioSinkID(dump)
	if err != nil {
		return -1, fmt.Errorf("acquiring ID: %w", err)
	}

	volume, err := getVolume(defaultAudioSinkID)
	if err != nil {
		return -1, fmt.Errorf("acquiring volume: %w", err)
	}

	return int(math.Round(volume * 100)), nil
}

func (c Client) GetDevice() (string, error) {
	dump, err := getPipewireDump()
	if err != nil {
		return "", fmt.Errorf("acquiring pipewire dump: %w", err)
	}

	defaultAudioSinkName, err := objects.GetDefaultAudioSinkName(dump)
	if err != nil {
		return "", fmt.Errorf("getting default audio sink name: %w", err)
	}

	if bluetoothRe.Match([]byte(defaultAudioSinkName)) {
		return DeviceTypeHeadphones, nil
	}

	return DeviceTypeSpeaker, nil
}
