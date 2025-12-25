package pipewire

import "fmt"

const (
	typePipeWireMetadata        = "PipeWire:Interface:Metadata"
	typePipeWireInterfaceNode   = "PipeWire:Interface:Node"
	metadataKeyDefaultAudioSink = "default.audio.sink"
)

type pwDumpResponseObjectMetadataValue struct {
	Name string `json:"name"`
}

type pwDumpResponseObjectMetadata struct {
	Key   string                            `json:"key"`
	Value pwDumpResponseObjectMetadataValue `json:"value"`
}

type pwDumpResponseObjectInfoParams struct {
	Props []map[string]any `json:"Props"`
}

type pwDumpResponseObjectInfo struct {
	NodeName string                         `json:"node.name"`
	Params   pwDumpResponseObjectInfoParams `json:"params"`
}

type pwDumpResponseObject struct {
	Type     string                         `json:"type"`
	Metadata []pwDumpResponseObjectMetadata `json:"metadata"`
	Info     pwDumpResponseObjectInfo       `json:"info"`
}

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

func getVolumeProp(props []map[string]any) (int, error) {
	relevantPropIndex := -1

	for index, prop := range props {
		if _, ok := prop["volume"]; ok {
			relevantPropIndex = index

			break
		}
	}

	if relevantPropIndex == -1 {
		return -1, fmt.Errorf("couldn't find %s prop", "volume")
	}

	floatVolume, ok := props[relevantPropIndex]["volume"].(float64)
	if !ok {
		return -1, fmt.Errorf("invalid %s data", "volume")
	}

	return int(floatVolume), nil
}

func getNodeByName(objects []pwDumpResponseObject, name string) (pwDumpResponseObject, error) {
	for _, obj := range objects {
		if obj.Info.NodeName == name {
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
