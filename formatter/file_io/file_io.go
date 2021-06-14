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
			fmt.Println("Failed to create temp directory")
			panic(err)
		}
	}
	filePath := path + "/temp." + ext
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create temp file")
		panic(err)
	}

	byteString := []byte(code)

	err = ioutil.WriteFile(filePath, byteString, 0700)
	if err != nil {
		fmt.Println("Failed to write temp file")
		panic(err)
	}
	defer f.Close()

}

func FormatFile(filePath string, command string, args []string) {
	_, err := exec.Command(command, args...).Output()
	if err != nil {
		fmt.Println("Failed to format temp file")
		panic(err)
	}
}
