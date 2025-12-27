package pipewire

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const (
	typePipeWireMetadata = "PipeWire:Interface:Metadata"

	metadataKeyDefaultAudioSink = "default.audio.sink"
)

func getDefaultAudioSinkName(objects []pwDumpResponseObject) (string, error) {
	metadataObjects := filterType(objects, typePipeWireMetadata)

	for _, obj := range metadataObjects {
		metadataIndex := metadataKeyIndex(obj, metadataKeyDefaultAudioSink)

		if metadataIndex != -1 {
			return obj.Metadata[metadataIndex].Value.Name, nil // id might be located in props, if not name + new search?,
		}
	}

	return "", fmt.Errorf("couldn't find %s", metadataKeyDefaultAudioSink)
}

func getProp(key string, props []map[string]any) (any, bool) {
	for _, prop := range props {
		if _, ok := prop[key]; ok {
			return prop[key], true
		}
	}

	return nil, false
}

func getVolumeProp(props []map[string]any) (float64, error) {
	prop, ok := getProp("volume", props)
	if !ok {
		return -1, fmt.Errorf("couldn't find %s prop", "volume")
	}

	floatVolume, ok := prop.(float64)
	if !ok {
		return -1, fmt.Errorf("invalid %s data", "volume")
	}

	return floatVolume, nil
}

func getNodeByName(objects []pwDumpResponseObject, name string) (pwDumpResponseObject, error) {
	for _, obj := range objects {
		nodeName, ok := obj.Info.Props["node.name"]
		if !ok {
			continue
		}

		if nodeName == name {
			return obj, nil
		}
	}

	return pwDumpResponseObject{}, fmt.Errorf("couldn't find node with name %s", name)
}

func filterType(objects []pwDumpResponseObject, objectType string) []pwDumpResponseObject {
	result := make([]pwDumpResponseObject, 0)

	for _, obj := range objects {
		if obj.Type == objectType {
			result = append(result, obj)
		}
	}

	return result
}

func metadataKeyIndex(obj pwDumpResponseObject, key string) int {
	for index, item := range obj.Metadata {
		if item.Key == key {
			return index
		}
	}

	return -1
}

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
