package battery

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const defaultBatteryPath = "/sys/class/power_supply/BAT0"

var defaultBatteryCapacityPath = path.Join(defaultBatteryPath, "capacity")
var defaultBatteryStatusPath = path.Join(defaultBatteryPath, "status")

func (b Status) String() string {
	icon := ""

	discharging := strings.Contains(strings.ToLower(b.Status), "discharging")

	if discharging {
		icon = "🔋"
	} else {
		icon = "⚡"
	}

	return fmt.Sprintf("%s%s%s", icon, b.Capacity, "%")
}

func GetStatus() (Status, error) {
	rawCapacity, err := os.ReadFile(defaultBatteryCapacityPath)
	if err != nil {
		return Status{}, fmt.Errorf("reading capacity: %w", err)
	}

	rawStatus, err := os.ReadFile(defaultBatteryStatusPath)
	if err != nil {
		return Status{}, fmt.Errorf("reading status: %w", err)
	}

	return Status{
		Capacity: strings.TrimSpace(string(rawCapacity)),
		Status:   string(rawStatus),
	}, nil
}
