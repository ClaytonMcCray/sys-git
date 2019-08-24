package sys_git

type worktree string

func Worktree(path string) worktree {
	return worktree(path)
}
