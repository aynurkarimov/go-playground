package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	fsAction := selectFsAction()

	switch fsAction {
	case "Create":
		fileName := selectFileName()
		createTextFile(fileName, ".txt")

	default:
		fileName := selectFileName()
		createTextFile(fileName, ".txt")
	}
}

var selectFsActionQuestion = &survey.Question{
	Name: "fsAction",
	Prompt: &survey.Select{
		Message: "Choose an action with a file:",
		Options: []string{"Create", "Delete", "Move"},
		Default: "Create",
	},
}

func selectFsAction() string {
	var fsActionAnswer string

	err := survey.AskOne(selectFsActionQuestion.Prompt, &fsActionAnswer)

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(0)
	}

	return fsActionAnswer
}

var fileNameQuestion = &survey.Question{
	Name: "fileName",
	Prompt: &survey.Input{
		Message: "Name a file that should be created",
	},
	Validate:  survey.Required,
	Transform: survey.ToLower,
}

func selectFileName() string {
	var fileName string

	err := survey.AskOne(fileNameQuestion.Prompt, &fileName)

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(0)
	}

	// Poor validation
	fileNameWithoutExtension := strings.Split(fileName, ".")[0]

	return fileNameWithoutExtension
}

func createTextFile(fileName string, extension string) error {
	file, err := os.Create(fileName + extension)

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(0)
	}

	defer file.Close()

	return nil
}
