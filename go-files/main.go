package main

import (
	"fe.com/go/files/fileManagement"
	"fmt"
	"os"
)

func main() {
	folderPath, _ := os.Getwd()
	filePath := folderPath + "/data/text.txt"
	text, err := fileUtils.ReadTextFile(folderPath + "/data/text.txt")
	if err == nil {
		fmt.Println(text)
		newContent := fmt.Sprintf("Original is: %v\n Double original is: %v%v", text, text, text)
		fileUtils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("ERROR PANIC!! %v", err)
	}
}
