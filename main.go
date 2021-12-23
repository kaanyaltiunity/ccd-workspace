package main

import "ccdWorkspace/services"

func main() {
	gitService := services.NewGitService()
	gitService.SetPreCommitHook()
}
