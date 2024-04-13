package utils

import (
	"io"
	"os"
)

func ReadFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	//defer file.Close()

	return file, err
}

func ReadFileContent(path string) ([]byte, error) {
	file, err := ReadFile(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
