package update

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s Status) String() string {
	return fmt.Sprintf("📦 %d", s.Packages)
}

func GetStatus() (Status, error) {
	output, _ := exec.Command("checkupdates").Output()

	return Status{
		Packages: strings.TrimSpace(string(output)),
	}, nil
}
