package services

import (
	"ccdWorkspace/utils"
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

const (
	commitRegex               = `(docs|feature|chore|fix)\((CCS)\-([0-9]*)\)\: ([a-zA-Z0-9_,. ]+$)|Merge branch`
	commitMessageHookFileName = "commit-msg"
)

type GitService interface {
	SetCommitMessageHook()
}

type gitService struct {
	hookTemplates HookTempaltes
}

type HookTempaltes struct {
	CommitMessageTemplate string
}

type CommitMessageTemplateVars struct {
	NodePath       string
	CommitRegex    string
	ColorModifiers utils.TextColorTags
}

func NewGitService(hookTemplates HookTempaltes) GitService {
	return &gitService{
		hookTemplates: hookTemplates,
	}
}

func (g *gitService) SetCommitMessageHook() {
	cmd := exec.Command("which", "node")
	nodePath, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(fmt.Sprintf("could not find \"node\", please install \"node\": %s", err))
	}

	commitMessageTemplate, err := template.New("Commit Message Js").Parse(g.hookTemplates.CommitMessageTemplate)
	if err != nil {
		log.Fatalln(fmt.Sprintf("ran into an error while parsing the commit message template: %s", err))
	}

	templateVars := CommitMessageTemplateVars{
		NodePath:       string(nodePath),
		CommitRegex:    commitRegex,
		ColorModifiers: utils.ColorTags,
	}

	err = g.writeToFile(commitMessageTemplate, templateVars, commitMessageHookFileName)
	if err != nil {
		log.Fatalln(err)
	}
}

func (g *gitService) writeToFile(template *template.Template, templateVars interface{}, fileName string) error {
	currentDir, _ := exec.Command("pwd").CombinedOutput()

	log.Printf("%screating %s in %s%s\n", utils.ColorTags.Foreground.Cyan, fileName, currentDir, utils.ColorTags.Modifiers.Reset)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("%serror while creating file\nFile Name: %s\nError: %s%s", utils.ColorTags.Foreground.Red, fileName, err, utils.ColorTags.Modifiers.Reset)
	}
	defer file.Close()

	err = template.Execute(file, templateVars)
	if err != nil {
		return fmt.Errorf("error while executing template for %s\nError: %s", fileName, err)
	}

	return nil
}

// func (g *gitService) getChildDirs()
