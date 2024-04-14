package repositories

import (
	"context"
	"time"

	"github.com/MorrisMorrison/retfig/database"
	"github.com/MorrisMorrison/retfig/models"
	uuid "github.com/satori/go.uuid"
)

const (
	QUERY_GET_PRESENTS_BY_EVENT_ID string = "SELECT * FROM presents WHERE event_id = ?"
)

type PresentRepository struct {
	dbConn *database.Connection
}

func NewPresentRepository(dbConn *database.Connection) *PresentRepository {
	return &PresentRepository{dbConn: dbConn}
}

func (repository *PresentRepository) GetPresentsByEventId(eventId uuid.UUID) ([]models.Present, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var presents []models.Present
	rows, queryErr := tx.QueryContext(ctx, QUERY_GET_PRESENTS_BY_EVENT_ID, eventId)
	if queryErr != nil {
		return nil, queryErr
	}

	for rows.Next() {
		var e models.Present
		err := rows.Scan(&e.Id, &e.EventId, &e.Creator, &e.Name, &e.Link, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}

		presents = append(presents, e)
	}

	defer rows.Close()

	return presents, tx.Commit()
}
