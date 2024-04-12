package migrations

import (
	"fmt"
	"log"
	"strings"

	"github.com/MorrisMorrison/retfig/database"
	"github.com/MorrisMorrison/retfig/utils"
)

func InitializeDatabase(connection *database.Connection) {
	statements, err := utils.ReadSqlFile("initialize-database.sql")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			log.Fatal("Could not find any sql statement in initialize-database.sql")
			break
		}

		_, err = connection.Database.Exec(statement)
		if err != nil {
			log.Fatalf("Failed to execute statement %q: %v", statement, err)
			break
		}
	}

	fmt.Println("SQL script executed successfully")
}
