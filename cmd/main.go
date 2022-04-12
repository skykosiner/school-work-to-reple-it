package main

import (
	"os"

	"github.com/yonikosiner/school-work-to-reple/pkg/work"
	"github.com/yonikosiner/school-work-to-reple/utils"
)

func main() {
	var w *work.Work
	home := os.Getenv("HOME")

	path := utils.AskForInput("What is the path for the project? (the path starts from ~/school-python/)")
	path = home + "/school-python/" + path[:len(path)-1]
	w.Run(path)
}
