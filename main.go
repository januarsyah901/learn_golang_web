package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo_new.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println("Version:", version)

	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	dirEntries, err := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("Directory:", entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
