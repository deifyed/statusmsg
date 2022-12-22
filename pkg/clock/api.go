package clock

import (
	"time"
)

func DTG() string {
	return time.Now().Format("021504")
}
