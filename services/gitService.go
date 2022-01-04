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
		log.Fatalln(fmt.Sprintf("%scould not find \"node\", please install \"node\": %s%s", utils.ColorTags.Foreground.Red, err, utils.ColorTags.Modifiers.Reset))
	}

	commitMessageTemplate, err := template.New("Commit Message Js").Parse(g.hookTemplates.CommitMessageTemplate)
	if err != nil {
		log.Fatalln(fmt.Sprintf("%sran into an error while parsing the commit message template: %s%s", utils.ColorTags.Foreground.Red, err, utils.ColorTags.Modifiers.Reset))
	}

	templateVars := CommitMessageTemplateVars{
		NodePath:       string(nodePath),
		CommitRegex:    commitRegex,
		ColorModifiers: utils.ColorTags,
	}
	childDirs, err := g.getChildDirs()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(childDirs)

	for _, childDir := range childDirs {
		os.Chdir(fmt.Sprintf("./%s", childDir))
		currentWorkingDir, err := os.Getwd()
		if err != nil {
			log.Fatalln("err")
		}
		log.Printf("%scurrent working dir%s %s", utils.ColorTags.Foreground.Cyan, utils.ColorTags.Modifiers.Reset, currentWorkingDir)
		subDirs, err := g.getChildDirs()
		if err != nil {
			log.Fatalln(err)
		}
		for _, subDir := range subDirs {
			if subDir == ".git" {
				os.Chdir(".git/hooks")
				currentWorkingDir, err = os.Getwd()
				if err != nil {
					log.Fatalln(err)
				}
				log.Printf("%scurrent working dir%s %s", utils.ColorTags.Foreground.Cyan, utils.ColorTags.Modifiers.Reset, currentWorkingDir)
				g.writeToFile(commitMessageTemplate, templateVars, commitMessageHookFileName)
				os.Chdir("../../")
			}
		}
		os.Chdir("../")
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func (g *gitService) writeToFile(template *template.Template, templateVars interface{}, fileName string) error {
	currentDir, _ := exec.Command("pwd").CombinedOutput()

	log.Printf("%screating %s in %s%s", utils.ColorTags.Foreground.Cyan, fileName, currentDir, utils.ColorTags.Modifiers.Reset)
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

func (g *gitService) getChildDirs() ([]string, error) {
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
