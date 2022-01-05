package main

import (
	"ccdWorkspace/cmd"
	"ccdWorkspace/services"
	_ "embed"
)

//go:embed templates/commit-msg.gotmpl
var commitMessageTemplate string

func main() {
	hookTemplates := services.HookTempaltes{
		CommitMessageTemplate: commitMessageTemplate,
	}
	gitHookService := services.NewGitHookService(hookTemplates)
	setGitHooksCmd := cmd.NewSetGitHooks(gitHookService)
	cmd.RootCmd.AddCommand(setGitHooksCmd)

	commitService := services.NewCommitValidatorService()
	validateCommitCmd := cmd.NewCommitValidator(commitService)
	cmd.RootCmd.AddCommand(validateCommitCmd)

	cmd.Execute()
}
