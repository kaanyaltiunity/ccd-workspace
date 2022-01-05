package services

import (
	"ccdWorkspace/utils"
	"regexp"
)

type CommitValidatorService interface {
	ValidateCommit(string) (bool, error)
}

type commitValidator struct {
}

func NewCommitValidatorService() CommitValidatorService {
	return &commitValidator{}
}

func (c *commitValidator) ValidateCommit(commitMessage string) (bool, error) {
	matched, err := regexp.MatchString(utils.CommitMessageReg, commitMessage)
	if err != nil {
		return false, err
	}

	if !matched {
		return false, nil
	}

	return true, nil
}
