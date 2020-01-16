package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := "./_ioutil.txt"
	content := "write by ioutil"

	var fileModeRW os.FileMode = 0666
	if err := ioutil.WriteFile(file, []byte(content), fileModeRW); err == nil {
		if fileBytes, err := ioutil.ReadFile(file); err == nil {
			fmt.Println(string(fileBytes))
		}
	}
}
