package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	fsAction := selectFsAction()

	switch fsAction {
	case "Create":
		fileName := selectFileName()
		createTextFile(fileName, ".txt")

	case "Delete":
		files := currentDirectoryFiles(".txt")
		selectedPaths := selectPathsForDeletion(files)

		for _, path := range selectedPaths {
			os.Remove(path)
		}

		fmt.Println("Success")

	default:
		fileName := selectFileName()
		createTextFile(fileName, ".txt")
	}
}

func selectFsAction() string {
	var fsActionAnswer string
	var selectFsActionQuestion = &survey.Question{
		Name: "fsAction",
		Prompt: &survey.Select{
			Message: "Choose an action with a file:",
			Options: []string{"Create", "Delete", "Move"},
			Default: "Create",
		},
	}

	err := survey.AskOne(selectFsActionQuestion.Prompt, &fsActionAnswer)

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(0)
	}

	return fsActionAnswer
}

func selectFileName() string {
	var fileName string
	var fileNameQuestion = &survey.Question{
		Name: "fileName",
		Prompt: &survey.Input{
			Message: "Name a file that should be created",
		},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	}

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

func currentDirectoryFiles(extension string) []string {
	var files []string

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(info.Name(), extension) {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(0)
	}

	return files
}

func selectPathsForDeletion(suggestions []string) []string {
	var paths []string
	var pathsForDeletionQuestion = &survey.MultiSelect{
		Message: "What files do you want to delete?",
		Options: suggestions,
	}

	err := survey.AskOne(pathsForDeletionQuestion, &paths)

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(0)
	}

	if len(paths) == 0 {
		fmt.Println(errors.New("please select at least one file"))

		os.Exit(0)
	}

	return paths
}
