package utils

import (
	"strings"
)

func ReadSqlFile(path string) ([]string, error) {
	content, err := ReadFileContent(path)
	if err != nil {
		return nil, err
	}

	sqlContent := string(content)

	statements := strings.Split(sqlContent, ";")
	return statements, nil
}
