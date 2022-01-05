package cmd

import (
	"ccdWorkspace/services"
	"ccdWorkspace/utils"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	validateCommitFlagName = "message"
	validateCommitShort    = "m"
)

var commitMessage string

type validateCommit struct {
	CommitValidatorService services.CommitValidatorService
}

func NewCommitValidator(cs services.CommitValidatorService) *cobra.Command {
	vc := &validateCommit{
		CommitValidatorService: cs,
	}
	cmd := &cobra.Command{
		Use:   "validateCommit",
		Short: "Validates commit message",
		RunE:  vc.ValidateCommit,
	}
	cmd.Flags().StringVarP(&commitMessage, validateCommitFlagName, validateCommitShort, "", "Commit message to validate")
	cmd.MarkFlagRequired(validateCommitFlagName)

	return cmd
}

func (v *validateCommit) ValidateCommit(cmd *cobra.Command, args []string) error {
	valid, err := v.CommitValidatorService.ValidateCommit(commitMessage)
	if err != nil {
		return err
	}

	if !valid {
		log.Printf(fmt.Sprintf("%sInvalid commit message%s >>>>> %s <<<<<", utils.ColorTags.Foreground.Red, utils.ColorTags.Modifiers.Reset, commitMessage))
		os.Exit(1)
	}

	log.Printf(fmt.Sprintf("%sValid commit message%s", utils.ColorTags.Foreground.Cyan, utils.ColorTags.Modifiers.Reset))
	return nil
}
