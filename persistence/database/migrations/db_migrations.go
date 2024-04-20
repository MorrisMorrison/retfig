package migrations

import (
	"fmt"
	"log"
	"strings"

	"github.com/MorrisMorrison/retfig/persistence/database"
	"github.com/MorrisMorrison/retfig/utils"
)

func InitializeDatabase(connection *database.Connection) {
	statements, err := utils.ReadSqlFile("database/sql/initialize-database.sql")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}

		_, err = connection.Database.Exec(statement)
		if err != nil {
			log.Fatalf("Failed to execute statement %q: %v", statement, err)
			break
		}
	}

	fmt.Println("SQL script executed successfully")
}
