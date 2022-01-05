package services

import (
	"ccdWorkspace/utils"
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"
)

const (
	commitMessageHookName = "commit-msg"
)

type SetHookFunc func() error

type GitHookService interface {
	GetHookFunc(string) (SetHookFunc, error)
	SetApplyToSubDir(bool)
}

type gitHookService struct {
	hooksMap      map[string]SetHookFunc
	hookTemplates HookTempaltes
	applyToSubDir bool
}

type HookTempaltes struct {
	CommitMessageTemplate string
}

type CommitMessageTemplateVars struct {
	NodePath       string
	CommitRegex    string
	ColorModifiers utils.TextColorTags
}

func NewGitHookService(hookTemplates HookTempaltes) GitHookService {
	service := &gitHookService{
		hookTemplates: hookTemplates,
	}
	hooksMap := make(map[string]SetHookFunc)
	hooksMap[commitMessageHookName] = service.setCommitMessageHook

	service.hooksMap = hooksMap

	return service
}

func (g *gitHookService) GetHookFunc(hookName string) (SetHookFunc, error) {
	if hookFunc, ok := g.hooksMap[hookName]; ok {
		return hookFunc, nil
	}
	return nil, fmt.Errorf("%scould not find any git hook with name%s \"%s\"", utils.ColorTags.Foreground.Red, utils.ColorTags.Modifiers.Reset, hookName)
}

func (g *gitHookService) SetApplyToSubDir(applyToSubDir bool) {
	g.applyToSubDir = applyToSubDir
}

func (g *gitHookService) setCommitMessageHook() error {
	cmd := exec.Command("which", "node")
	nodePath, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Errorf("%scould not find \"node\", please install \"node\": %s%s", utils.ColorTags.Foreground.Red, err, utils.ColorTags.Modifiers.Reset)
	}

	commitMessageTemplate, err := template.New("Commit Message Js").Parse(g.hookTemplates.CommitMessageTemplate)
	if err != nil {
		fmt.Errorf("%sran into an error while parsing the commit message template: %s%s", utils.ColorTags.Foreground.Red, err, utils.ColorTags.Modifiers.Reset)
	}

	templateVars := CommitMessageTemplateVars{
		NodePath:       string(nodePath),
		CommitRegex:    utils.CommitMessageReg,
		ColorModifiers: utils.ColorTags,
	}
	childDirs, err := g.getChildDirs()
	if err != nil {
		return err
	}

	if g.applyToSubDir {
		for _, childDir := range childDirs {
			os.Chdir(fmt.Sprintf("./%s", childDir))
			currentWorkingDir, err := os.Getwd()
			if err != nil {
				return err
			}
			log.Printf("%scurrent working dir%s %s", utils.ColorTags.Foreground.Cyan, utils.ColorTags.Modifiers.Reset, currentWorkingDir)
			subDirs, err := g.getChildDirs()
			if err != nil {
				return err
			}
			for _, subDir := range subDirs {
				if subDir == ".git" {
					os.Chdir(".git/hooks")
					currentWorkingDir, err = os.Getwd()
					if err != nil {
						return err
					}
					log.Printf("%scurrent working dir%s %s", utils.ColorTags.Foreground.Cyan, utils.ColorTags.Modifiers.Reset, currentWorkingDir)
					g.writeToFile(commitMessageTemplate, templateVars, commitMessageHookName)
					os.Chdir("../../")
				}
			}
			os.Chdir("../")
		}
		return nil
	}
	isGit := false
	for _, childDir := range childDirs {
		if childDir == ".git" {
			isGit = true
		}
	}
	if !isGit {
		log.Printf("%scurrent directory is not a git repo%s", utils.ColorTags.Foreground.Cyan, utils.ColorTags.Modifiers.Reset)
		return nil
	}
	os.Chdir("./git/hooks")
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Printf("%scurrent working dir%s %s", utils.ColorTags.Foreground.Cyan, utils.ColorTags.Modifiers.Reset, currentWorkingDir)

	err = g.writeToFile(commitMessageTemplate, templateVars, commitMessageHookName)
	if err != nil {
		return err
	}

	return nil
}

func (g *gitHookService) writeToFile(template *template.Template, templateVars interface{}, fileName string) error {
	currentDir, _ := exec.Command("pwd").CombinedOutput()

	log.Printf("%screating %s in%s %s", utils.ColorTags.Foreground.Cyan, fileName, utils.ColorTags.Modifiers.Reset, currentDir)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("%serror while creating file\nFile Name: %s\nError: %s%s", utils.ColorTags.Foreground.Red, fileName, err, utils.ColorTags.Modifiers.Reset)
	}
	defer file.Close()

	log.Printf("%sexecuting template in %s%s", utils.ColorTags.Foreground.Cyan, fileName, utils.ColorTags.Modifiers.Reset)
	err = template.Execute(file, templateVars)
	if err != nil {
		return fmt.Errorf("%sran into an error while executing template for %s\nError: %s%s", utils.ColorTags.Foreground.Red, fileName, err, utils.ColorTags.Modifiers.Reset)
	}

	return nil
}

func (g *gitHookService) getChildDirs() ([]string, error) {
	var childDirs []string
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("%sran into an error while fetching contents of the current directory\nError: %s%s", utils.ColorTags.Foreground.Red, err, utils.ColorTags.Modifiers.Reset)
	}

	for _, info := range files {
		if info.IsDir() {
			childDirs = append(childDirs, info.Name())
		}
	}
	return childDirs, nil
}
