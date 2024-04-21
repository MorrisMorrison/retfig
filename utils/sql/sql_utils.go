package sql

import (
	"strings"

	"github.com/MorrisMorrison/retfig/utils/file"
)

func ReadSqlFile(path string) ([]string, error) {
	content, err := file.ReadFileContent(path)
	if err != nil {
		return nil, err
	}

	sqlContent := string(content)

	statements := strings.Split(sqlContent, ";")
	return statements, nil
}
