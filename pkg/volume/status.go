package volume

import "strings"

func (s Status) isMuted() bool {
	return strings.ToLower(s.Level) == "m"
}
