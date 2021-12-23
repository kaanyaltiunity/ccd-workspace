package main

import (
	"ccdWorkspace/services"
	_ "embed"
)

//go:embed templates/commit-msg.gotmpl
var commitMessageTemplate string

func main() {
	hookTemplates := services.HookTempaltes{
		CommitMessageTemplate: commitMessageTemplate,
	}
	gitService := services.NewGitService(hookTemplates)
	gitService.SetCommitMessageHook()
}
