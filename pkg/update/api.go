package update

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s Status) String() string {
	return fmt.Sprintf("📦%d", s.PackageCount)
}

func GetStatus() (Status, error) {
	output, _ := exec.Command("checkupdates").Output()

	outputAsString := strings.TrimSpace(string(output))

	var count int

	if outputAsString == "" {
		count = 0
	} else {
		count = len(strings.Split(outputAsString, "\n"))
	}

	return Status{
		PackageCount: count,
	}, nil
}
