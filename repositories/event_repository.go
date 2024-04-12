package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/MorrisMorrison/retfig/database"
	"github.com/MorrisMorrison/retfig/models"
)

type EventRepository struct {
	dbConn *database.Connection
}

func NewEventRepository(dbConn *database.Connection) *EventRepository {
	return &EventRepository{dbConn: dbConn}
}

func (repository *EventRepository) CreateEvent(event models.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := repository.dbConn.ExecuteInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, execError := tx.ExecContext(ctx, "INSERT INTO events VALUES (?, ?, ?)", event.Name, event.Owner.Id, event.Recipient.Id)
		if execError != nil {
			return execError
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (repository *EventRepository) GetEventById(id string) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var e models.Event
	queryErr := tx.QueryRowContext(ctx, "SELECT * FROM events WHERE id = ?", id).Scan(&e.Name, &e.Owner.Id, &e.Recipient.Id)
	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, queryErr
		}
	}

	return &e, tx.Commit()
}

func (repository *EventRepository) GetEventsByIds(ids []string) ([]models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var events []models.Event
	rows, queryErr := tx.QueryContext(ctx, "SELECT * FROM events WHERE id IN (?)", ids)
	if queryErr != nil {
		return nil, queryErr
	}

	for rows.Next() {
		var e models.Event
		err := rows.Scan(&e.Name, &e.Owner, &e.Recipient)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	defer rows.Close()

	return events, tx.Commit()
}
