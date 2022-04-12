package work

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/yonikosiner/school-work-to-reple/utils"
)

type Work struct {
	Path     string   `json:"path"`
	Contents []string `json:"contents"`
}

func (w *Work) Run(path string) {
	files := w.getFilesFromPath(path)

	for _, value := range files {
		contents := w.getFileContents(value)
		fmt.Println(contents)
	}
}

func (w *Work) getFileContents(path string) Work {
	b, err := os.Open(path)
	fmt.Println("getFileContents path", path)

	if err != nil {
		log.Fatal(fmt.Sprintf("getFileContents (open file) %s", err))
	}

	scanner := bufio.NewScanner(b)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	b.Close()

	var contents []string

	for _, value := range text {
		contents = append(contents, fmt.Sprintf("%s\n", value))
	}

	return Work{path, contents}
}

func (w *Work) getFilesFromPath(path string) []string {
	var filesList []string
	var dirFileList []string

	// God this looks ugly, but it works
    // TODO: ignore pdf files
	ignoreList := []string{"node_modules", ".git", "__pycache__", ".gitignore", "README.md"}
	for _, files := range utils.GetFilesFromPath(path) {
		if !utils.Contains(ignoreList, files.Name()) {
			if files.IsDir() {
				for _, filefiles := range utils.GetFilesFromPath(path) {
					dirFileList = append(dirFileList, fmt.Sprintf("%s/%s", path, filefiles.Name()))
				}
			}
			filesList = append(filesList, fmt.Sprintf("%s/%s", path, files.Name()))
		}
	}

	filesList = append(filesList, dirFileList...)
	return filesList
}
