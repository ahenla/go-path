package fileUtils

import "os"

func ReadTextFile(fileName string) (string, error) {
  content, err := os.ReadFile(fileName)
  if err != nil {
    //could't read the file 
    return "", err
  }else {
    //read successful
    return string(content), nil
  }
  
}
