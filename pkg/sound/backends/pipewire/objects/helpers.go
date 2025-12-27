package objects

import (
	"encoding/json"
	"fmt"
	"io"
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

func parseDump(dump io.Reader) ([]pwDumpResponseObject, error) {
	var response []pwDumpResponseObject

	err := json.NewDecoder(dump).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}

	return response, nil
}
