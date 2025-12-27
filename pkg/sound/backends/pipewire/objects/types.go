package objects

import (
	"encoding/json"
	"errors"
)

const (
	typePipeWireMetadata = "PipeWire:Interface:Metadata"

	metadataKeyDefaultAudioSink = "default.audio.sink"
)

type pwDumpResponseObjectMetadataValue struct {
	IsObject bool
	Name     string `json:"name"`

	IsInt      bool
	ValueAsInt int

	IsBool      bool
	ValueAsBool bool
}

func (v *pwDumpResponseObjectMetadataValue) UnmarshalJSON(data []byte) error {
	var obj struct {
		Name string `json:"name"`
	}

	if err := json.Unmarshal(data, &obj); err == nil && obj.Name != "" {
		v.Name = obj.Name
		v.IsObject = true

		return nil
	}

	var intValue int
	if err := json.Unmarshal(data, &intValue); err != nil {
		v.ValueAsInt = intValue
		v.IsInt = true

		return nil
	}

	var boolValue bool
	if err := json.Unmarshal(data, &boolValue); err != nil {
		v.ValueAsBool = boolValue
		v.IsBool = true

		return nil
	}

	return errors.New("value is neither struct, bool nor int")
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
	Props    map[string]any
}

type pwDumpResponseObject struct {
	ID       int                            `json:"id"`
	Type     string                         `json:"type"`
	Metadata []pwDumpResponseObjectMetadata `json:"metadata"`
	Info     pwDumpResponseObjectInfo       `json:"info"`
}
