package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	selectFsAction()
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
