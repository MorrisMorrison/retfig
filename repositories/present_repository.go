package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/MorrisMorrison/retfig/database"
	"github.com/MorrisMorrison/retfig/models"
	uuid "github.com/satori/go.uuid"
)

const (
	QUERY_CREATE_PRESENT           string = "INSERT INTO present (id, event_id, creator, name, link, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?)"
	QUERY_GET_PRESENTS_BY_EVENT_ID string = "SELECT * FROM present WHERE event_id = ?"
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

func (repository *PresentRepository) CreatePresent(present models.Present) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newUUID := uuid.NewV4()
	now := time.Now()

	err := repository.dbConn.ExecuteInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, execError := tx.ExecContext(ctx, QUERY_CREATE_PRESENT, newUUID.String(), present.EventId, present.Creator, present.Name, present.Link, now, now)
		if execError != nil {
			return execError
		}

		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil
}
