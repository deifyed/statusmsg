package pipewire

import (
	"fmt"
)

const (
	typePipeWireMetadata        = "PipeWire:Interface:Metadata"
	typePipeWireInterfaceNode   = "PipeWire:Interface:Node"
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

func getVolumeProp(props []map[string]any) (int, error) {
	prop, ok := getProp("volume", props)
	if !ok {
		return -1, fmt.Errorf("couldn't find %s prop", "volume")
	}

	floatVolume, ok := prop.(float64)
	if !ok {
		return -1, fmt.Errorf("invalid %s data", "volume")
	}

	return int(floatVolume), nil
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

/*
pw-dump | jq '.[] | select(.type == "PipeWire:Interface:Metadata")'
*/
