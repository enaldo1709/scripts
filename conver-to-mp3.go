package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)

	
	cd, err := os.ReadDir(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, file := range cd {
		
		//fmt.Printf("file entry -> %s - %s\n", file.Type(), file.Name())
		name, ext := getNameAndExtension(file)
		fmt.Printf("file: %s\text: %s\n", name, ext)
	}

}

func getNameAndExtension(f fs.DirEntry) (string, string) {
	if (f.IsDir()) {
		return f.Name(), ""
	}

	lastDot := strings.LastIndex(f.Name(), ".")
	return f.Name()[:lastDot], f.Name()[lastDot:]

}
