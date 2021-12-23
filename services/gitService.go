package services

import (
	"fmt"
	"log"
	"os/exec"
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
	fmt.Println(string(output))
}
