// Package countdown handles displaying days left until a date
package countdown

import (
	"fmt"
	"math"
	"time"
)

func Days(date string) (int, error) {
	target, err := time.ParseInLocation(time.DateOnly, date, time.Local)
	if err != nil {
		return 0, fmt.Errorf("parsing date: %w", err)
	}

	y, m, d := time.Now().Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	return int(math.Round(target.Sub(today).Hours() / 24)), nil
}
