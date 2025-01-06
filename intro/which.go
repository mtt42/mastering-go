package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Not enough arguments.")
		return
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

  fmt.Println("The following paths are examined:")
  for _, path := range pathSplit {
    fmt.Println(path)
  }

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)

		// Does the file exist?
		fileinfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileinfo.Mode()
		if !mode.IsRegular() {
			continue
		}

		if mode&0111 != 0 {
			fmt.Println(fullPath)
			return
		}

	}
}
