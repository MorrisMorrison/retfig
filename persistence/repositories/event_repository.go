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
	QUERY_CREATE_EVENT      string = "INSERT INTO event (id, name, recipient, createdBy, updatedBy, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?)"
	QUERY_GET_EVENT_BY_ID   string = "SELECT * FROM event WHERE id = ?"
	QUERY_GET_EVENTS_BY_IDS string = "SELECT * FROM event WHERE id in (?)"
)

type EventRepository struct {
	dbConn *database.Connection
}

func NewEventRepository(dbConn *database.Connection) *EventRepository {
	return &EventRepository{dbConn: dbConn}
}

func (repository *EventRepository) CreateEvent(event models.Event) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newUUID := uuid.NewV4()
	now := time.Now()

	err := repository.dbConn.ExecuteInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, execError := tx.ExecContext(ctx, QUERY_CREATE_EVENT, newUUID.String(), event.Name, event.Recipient, event.CreatedBy, event.UpdatedBy, now, now)
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

func (repository *EventRepository) GetEventById(id uuid.UUID) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var e models.Event
	queryErr := tx.QueryRowContext(ctx, QUERY_GET_EVENT_BY_ID, id).Scan(&e.Id, &e.Name, &e.Recipient, &e.CreatedBy, &e.UpdatedBy, &e.CreatedAt, &e.UpdatedAt)
	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, queryErr
		}
	}

	return &e, tx.Commit()
}

func (repository *EventRepository) GetEventsByIds(ids []string) ([]*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var events []*models.Event
	rows, queryErr := tx.QueryContext(ctx, QUERY_GET_EVENTS_BY_IDS, ids)
	if queryErr != nil {
		return nil, queryErr
	}

	for rows.Next() {
		var e models.Event
		err := rows.Scan(&e.Id, &e.Name, &e.Recipient, &e.CreatedBy, &e.UpdatedBy, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}

		events = append(events, &e)
	}

	defer rows.Close()

	return events, tx.Commit()
}
