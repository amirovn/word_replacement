package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
	"github.com/iancoleman/strcase"
)

var showLog = true

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Arguments")
		fmt.Println("search - required - (Search the string and replace file name and content)")
		fmt.Println("replace - required - (Replace the string from search)")
		fmt.Println("path - required - (Search and replace the path)")
		fmt.Println("--quiet - (Not show logs)")

		return
	}

	search := os.Args[1]
	replace := os.Args[2]

	filesNotExists(os.Args[3])

	files := getFilesPath(os.Args[3])

	if len(os.Args) > 4 && os.Args[4] == "--quiet" {
		showLog = false
	}

	replaceFileContent(files, search, replace)
}

func replaceFileContent(files []string, search string, replace string) {
	for _, file := range files {
		listCamels := transferCamels(search, replace)
		updateShowFileLog := 0

		for searchCamel, replaceCamel := range listCamels {

			read, err := ioutil.ReadFile(file)

			if err != nil {
				fmt.Println(err)
			}

			if strings.Contains(string(read), searchCamel) {
				contents := strings.ReplaceAll(string(read), searchCamel, replaceCamel)

				err = ioutil.WriteFile(file, []byte(contents), 0644)

				if err != nil {
					fmt.Println(err)
				} else if showLog == true {
					if updateShowFileLog == 0 {
						updateShowFileLog = 1
						fmt.Printf("Update file content: %v \n", file)
					}

					fmt.Printf(" from %v to %v\n", searchCamel, replaceCamel)
				}
			}
		}
	}

	replaceFileName(files, search, replace)
	replaceDirectoryName(search, replace)
}

func replaceFileName(files []string, search string, replace string) {
	for _, file := range files {

		listCamels := transferCamels(search, replace)

		for searchCamel, replaceCamel := range listCamels {
			fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))

			if strings.Contains(fileName, searchCamel) {
				fileNewPath := strings.Replace(file, searchCamel, replaceCamel, 1)

				defer os.Rename(file, fileNewPath)

				if showLog == true {
					fmt.Printf("Rename file name:\n from %v \n to %v \n", file, fileNewPath)
				}
			}
		}
	}
}

func replaceDirectoryName(search string, replace string) {
	directories := getDirectories(os.Args[3])

	for _, directory := range directories {

		listCamels := transferCamels(search, replace)

		for searchCamel, replaceCamel := range listCamels {
			if strings.Contains(directory, searchCamel) {
				directoryNewPath := strings.Replace(directory, searchCamel, replaceCamel, 1)

				defer os.Rename(directory, directoryNewPath)

				if showLog == true {
					fmt.Printf("Rename directory name:\n from %v \n to %v \n", directory, directoryNewPath)
				}
			}
		}
	}
}

func getFilesPath(pathName string) []string {
	files := make([]string, 0)
	err := filepath.Walk(pathName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		if info.IsDir() {
			return nil
		}

		if _, ok := filetype.Types[strings.Trim(filepath.Ext(path), ".")]; ok {
			return nil
		}

		files = append(files, path)

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return files
}

func getDirectories(pathName string) []string {
	directories := make([]string, 0)

	err := filepath.Walk(pathName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		if info.IsDir() {
			directories = append(directories, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return directories
}

func transferCamels(search string, replace string) map[string]string {
	list := make(map[string]string)

	list[strcase.ToCamel(search)] = strcase.ToCamel(replace)
	list[strcase.ToLowerCamel(search)] = strcase.ToLowerCamel(replace)
	list[strcase.ToKebab(search)] = strcase.ToKebab(replace)
	list[strcase.ToScreamingKebab(search)] = strcase.ToScreamingKebab(replace)
	list[strcase.ToSnake(search)] = strcase.ToSnake(replace)
	list[strcase.ToScreamingSnake(search)] = strcase.ToScreamingSnake(replace)
	list[strings.ToUpper(search)] = strings.ToUpper(replace)
	list[strings.ToLower(search)] = strings.ToLower(replace)

	return list
}

func filesNotExists(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		panic(err)
	}
}
