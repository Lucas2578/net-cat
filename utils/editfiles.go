package utils

import (
	"fmt"
	"os"
)

func WriteFile(fileNameLogs string, message string) {
	message = regexRemoveAnsi.ReplaceAllString(message, "")
	file, err := os.OpenFile(url+fileNameLogs, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error when open \"%v\" file\nError: \"%v\"\n", fileNameLogs, err)
		return
	}

	if _, err := file.WriteString(message); err != nil {
		fmt.Printf("Error when write on \"%v\" file\nError: \"%v\"\n", fileNameLogs, err)
		return
	}
	defer file.Close()
}
