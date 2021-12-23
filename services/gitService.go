package services

import (
	"fmt"
	"log"
	"os/exec"
	"text/template"
)

const (
	commitRegex = `(docs|feature|chore)\((CCD)\-([0-9]*)\)\: ([a-zA-Z0-9_,. ]+$)|Merge branch`
)

type GitService interface {
	SetPreCommitHook()
}

type gitService struct {
}

func NewGitService() GitService {
	return &gitService{}
}

func (g *gitService) SetPreCommitHook() {
	cmd := exec.Command("which", "node")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(fmt.Sprintf("could not find \"node\", please install \"node\": %s", err))
	}
	//go:embed "commit-msg.gotmpl"
	var validatorTemplate string
	validatorTemplate, err = template.New("Commit Message JS")
	fmt.Println(string(output))
}
