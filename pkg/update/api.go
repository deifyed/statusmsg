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
	output, err := exec.Command("checkupdates").Output()
	if err != nil {
		return Status{}, err
	}

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
