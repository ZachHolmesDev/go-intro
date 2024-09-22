package main

import (
	"files/fileUtils"
	"fmt"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	content, err := fileUtils.ReadFile(filePath)

	if err != nil {
		fmt.Printf("panic !! %v\n", err)
	} else {
		fmt.Println(content)

		newContent := fmt.Sprintf("Origional %v \n Double origional %v%v", content, content, content)
		// newContent := "Origional : "+ content

		fileUtils.WriteToFile(filePath+".output.txt", newContent)
	}
}
