package pipewire

import (
	"bytes"
	"encoding/json"
	"errors"
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
}

type pwDumpResponseObject struct {
	Type     string                         `json:"type"`
	Metadata []pwDumpResponseObjectMetadata `json:"metadata"`
	Info     pwDumpResponseObjectInfo       `json:"info"`
}
