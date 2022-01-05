package cmd

import (
	"ccdWorkspace/services"

	"github.com/spf13/cobra"
)

const (
	hookFlagName = "hook"
	allFlagName  = "all"
)

var (
	setCommitMessageFlag string
	setHookForAll        bool
)

type setGitHooks struct {
	gitHookService services.GitHookService
}

func NewSetGitHooks(gitHookService services.GitHookService) *cobra.Command {

	sgh := &setGitHooks{
		gitHookService: gitHookService,
	}
	cmd := &cobra.Command{
		Use:   "setGitHooks",
		Short: "Sets git hooks",
		Long:  "Sets git hooks in all the sub directories that are git repositories. These hooks are placed in \".git/hooks/<HOOK NAME>\" in each repo",
		RunE:  sgh.SetGitHooksCmd,
	}

	cmd.Flags().StringVarP(&setCommitMessageFlag, hookFlagName, "", "", "Git hook to set")
	cmd.MarkFlagRequired(hookFlagName)

	cmd.Flags().BoolVarP(&setHookForAll, allFlagName, "", false, "Determines whether to apply the commit hook to all sub directories or only in the current working directory")
	return cmd
}

func (s *setGitHooks) SetGitHooksCmd(cmd *cobra.Command, args []string) error {
	s.gitHookService.SetApplyToSubDir(setHookForAll)
	hookFunc, err := s.gitHookService.GetHookFunc(setCommitMessageFlag)
	if err != nil {
		return err
	}

	err = hookFunc()
	if err != nil {
		return err
	}
	return nil
}
