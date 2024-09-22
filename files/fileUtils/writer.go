package fileUtils

import "os"

func WriteToFile(fileName string, content string) error {
	return os.WriteFile(fileName, []byte(content), 0644)
}
