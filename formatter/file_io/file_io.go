package file_io

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type Input struct {
	Lang      string
	Formatter string
	Code      string
	Args      []string
}

func WriteTempFile(ext string, code string) {
	path := "./tmp"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0700)
		if err != nil {
			fmt.Printf("Failed to create temp directory, %s", err)
		}
	}
	filePath := path + "/temp." + ext
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Failed to create temp file, %s", err)
	}

	byteString := []byte(code)

	err = ioutil.WriteFile(filePath, byteString, 0700)
	if err != nil {
		fmt.Printf("Failed to write temp file, %s", err)
	}
	defer f.Close()

}

func FormatFile(filePath string, command string, args []string) {
	_, err := exec.Command(command, args...).Output()
	if err != nil {
		fmt.Printf("Failed to format temp file, %s", err)
	}
}
