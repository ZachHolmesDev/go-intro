package fileUtils

import "os"

func ReadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
		// cant read file
		// return "" // lazy version
		// panic("hahahaha")
	} else {
		// read file sucess
		return string(content), nil
	}
}
