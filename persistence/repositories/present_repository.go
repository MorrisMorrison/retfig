package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/MorrisMorrison/retfig/persistence/database"
	"github.com/MorrisMorrison/retfig/persistence/models"
	uuid "github.com/satori/go.uuid"
)

const (
	QUERY_CREATE_PRESENT           string = "INSERT INTO present (id, eventId, name, link, createdBy, updatedBy, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	QUERY_GET_PRESENTS_BY_EVENT_ID string = "SELECT * FROM present WHERE eventId = ?"
	QUERY_GET_PRESENT_BY_ID        string = "SELECT * FROM present WHERE id = ?"
)

type PresentRepository struct {
	dbConn *database.Connection
}

func NewPresentRepository(dbConn *database.Connection) *PresentRepository {
	return &PresentRepository{dbConn: dbConn}
}

func (repository *PresentRepository) GetPresentsByEventId(eventId uuid.UUID) ([]*models.Present, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var presents []*models.Present
	rows, queryErr := tx.QueryContext(ctx, QUERY_GET_PRESENTS_BY_EVENT_ID, eventId)
	if queryErr != nil {
		return nil, queryErr
	}

	for rows.Next() {
		var p models.Present
		err := rows.Scan(&p.Id, &p.EventId, &p.Name, &p.Link, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}

		presents = append(presents, &p)
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
		_, execError := tx.ExecContext(ctx, QUERY_CREATE_PRESENT, newUUID.String(), present.EventId, present.Name, present.Link, present.CreatedBy, present.UpdatedBy, now, now)
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

func (repository *PresentRepository) GetPresentById(id uuid.UUID) (*models.Present, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var p models.Present
	queryErr := tx.QueryRowContext(ctx, QUERY_GET_PRESENT_BY_ID, id).Scan(&p.Id, &p.EventId, &p.Name, &p.Link, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, queryErr
		}
	}

	return &p, tx.Commit()
}
