package cmd

import (
	services "ccdWorkspace/services/git"

	"github.com/spf13/cobra"
)

const (
	flagName = "hook"
)

var (
	setCommitMessageFlag string
)

type SetGitHooks interface {
	SetGitHooksCmd(*cobra.Command, []string) error
}

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

	cmd.Flags().StringVarP(&setCommitMessageFlag, flagName, "", "", "Git hook to set")
	cmd.MarkFlagRequired(flagName)
	return cmd
}

func (s *setGitHooks) SetGitHooksCmd(cmd *cobra.Command, args []string) error {
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
