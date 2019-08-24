package sys_git

import (
	"os/exec"
	"strings"
)

type worktree string

func Worktree(path string) worktree {
	return worktree(path)
}

func (w worktree) Status() (map[string]string, error) {
	cmd := exec.Command("git", "-C", string(w), "--porcelain")
	buffOut, err := cmd.Output()
	if err != nil {
		return map[string]string{}, err
	}

	bufferedStatus := string(buffOut)
	splitStatus := strings.Split(bufferedStatus, " ")
	status := make(map[string]string)
	for i := 0; i < len(splitStatus); i += 2 {
		status[splitStatus[i]] = splitStatus[i+1]
	}

	return status, nil
}
