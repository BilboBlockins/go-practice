package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//go can be used to clean up dir structure in a file system
func main() {
	//absolute filepath - dot denotes the current dir
	root, _ := filepath.Abs(".")
	fmt.Println("Processing Path", root)

	err := filepath.Walk(root, processPath)
	if err != nil {
		fmt.Println("error:", err)
	}
}

//args and types are important
func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	
	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
	}
	return nil
}
