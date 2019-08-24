package sys_git

import (
	"os/exec"
	"strings"
)

// TODO document possible file statuses
const (
	Untracked = "??"
)

func (w worktree) Status() (map[string]string, error) {
	buffOut, err := exec.Command("git", "-C", string(w), "status", "--porcelain", "-z").Output()
	if err != nil {
		return map[string]string{}, err
	}

	strStatus := string(buffOut)
	splitStatus := strings.Split(strStatus, "\000") // "-z" causes split on unicode nul
	status := make(map[string]string)
	for _, pair := range splitStatus {
		if len([]rune(pair)) >= 4 {
			status[pair[3:]] = pair[0:2] // pair looks like "XY path/to/file" per git-scm docs
		}
	}

	return status, nil
}
