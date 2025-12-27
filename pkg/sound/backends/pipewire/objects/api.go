// Package objects handles reading from pw-dump
package objects

import (
	"fmt"
	"io"
)

func GetDefaultAudioSinkID(dump io.Reader) (int, error) {
	objects, err := parseDump(dump)
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

	return defaultAudioSinkNode.ID, nil
}

func GetDefaultAudioSinkName(dump io.Reader) (string, error) {
	objects, err := parseDump(dump)
	if err != nil {
		return "", fmt.Errorf("getting objects: %w", err)
	}

	defaultAudioSinkName, err := getDefaultAudioSinkName(objects)
	if err != nil {
		return "", fmt.Errorf("getting default audio sink name: %w", err)
	}

	return defaultAudioSinkName, nil
}
