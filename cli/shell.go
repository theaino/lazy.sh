package cli

import (
	"fmt"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func fetchShell() (shell string, err error) {
	user, err := user.Current()
	if err != nil {
		return
	}

	out, err := exec.Command("getent", "passwd", user.Uid).Output()
	if err != nil {
		err = fmt.Errorf("Failed to fetch shell for user %s: %w", user.Name, err)
		return
	}

	ent := strings.Split(strings.Trim(string(out), "\n"), ":")
	shell = filepath.Base(ent[6])
	return
}
