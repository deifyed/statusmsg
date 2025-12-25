package pipewire

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/pw-dump.json
var pwDump string

func TestParsing(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name string
	}{
		"Should test name": {},
	}

	for name := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var response []pwDumpResponseObject
			err := json.NewDecoder(strings.NewReader(pwDump)).Decode(&response)
			assert.NoError(t, err)
		})
	}
}
