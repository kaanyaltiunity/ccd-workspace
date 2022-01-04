package main

import (
	"ccdWorkspace/cmd"
	services "ccdWorkspace/services/git"
	_ "embed"
)

//go:embed templates/commit-msg.gotmpl
var commitMessageTemplate string

func main() {
	hookTemplates := services.HookTempaltes{
		CommitMessageTemplate: commitMessageTemplate,
	}
	gitHookService := services.NewGitHookService(hookTemplates)
	gitHooksCmd := cmd.NewSetGitHooks(gitHookService)
	cmd.RootCmd.AddCommand(gitHooksCmd)

	cmd.Execute()
}
