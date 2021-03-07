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

	count := len(strings.Split(strings.TrimSpace(string(output)), "\n"))

	return Status{
		PackageCount: count,
	}, nil
}
