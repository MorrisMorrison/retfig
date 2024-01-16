package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/MorrisMorrison/retfig/database"
	"github.com/MorrisMorrison/retfig/models"
)

func CreateEvent(event models.Event) error {
	err := database.ExecuteInTransaction(func(ctx context.Context, tx *sql.Tx) error {
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

func GetEventById(id string) (*models.Event, error) {
	db := database.GetDbConnection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := db.BeginTx(ctx, nil)
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

func GetEventsByIds(ids []string) ([]models.Event, error) {
	db := database.GetDbConnection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := db.BeginTx(ctx, nil)
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
