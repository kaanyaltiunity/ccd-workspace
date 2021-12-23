package services

import "fmt"

type GitService interface {
	SetPreCommitHook()
}

type gitService struct {
}

func NewGitService() GitService {
	return &gitService{}
}

func (g *gitService) SetPreCommitHook() {
	fmt.Println("D")
}
