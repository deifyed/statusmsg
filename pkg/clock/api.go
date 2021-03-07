package clock

import (
	"fmt"
	"time"
)

func (s Status) String() string {
	return fmt.Sprintf("🕙%02d%02d%02d", s.Time.Day(), s.Time.Hour(), s.Time.Minute())
}

func GetStatus() Status {
	return Status{
		Time: time.Now(),
	}
}
