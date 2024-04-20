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
	QUERY_GET_PARTICIPANTS_BY_EVENT_ID         string = "SELECT * FROM event_participant WHERE eventId = ?"
	QUERY_GET_PARTICIPANT_BY_NAME_AND_EVENT_ID string = "SELECT * FROM event_participant WHERE name=? AND eventId=?"
	QUERY_CREATE_PARTICIPANT                   string = "INSERT INTO event_participant (eventId, name, createdBy, updatedBy, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)"
)

type ParticipantRepository struct {
	dbConn *database.Connection
}

func NewParticipantRepository(dbConn *database.Connection) *ParticipantRepository {
	return &ParticipantRepository{dbConn: dbConn}
}

func (repository *ParticipantRepository) CreateParticipant(eventId string, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()

	err := repository.dbConn.ExecuteInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, execError := tx.ExecContext(ctx, QUERY_CREATE_PARTICIPANT, eventId, username, now, now)
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

func (repository *ParticipantRepository) GetParticipantsByEventId(eventId uuid.UUID) ([]*models.Participant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var participants []*models.Participant
	rows, queryErr := tx.QueryContext(ctx, QUERY_GET_PARTICIPANTS_BY_EVENT_ID, eventId)
	if queryErr != nil {
		return nil, queryErr
	}

	for rows.Next() {
		var p models.Participant
		err := rows.Scan(&p.EventId, &p.Name, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}

		participants = append(participants, &p)
	}

	defer rows.Close()

	return participants, tx.Commit()
}

func (repository *ParticipantRepository) GetParticipantByNameAndEventId(name string, eventId uuid.UUID) (*models.Participant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var e models.Participant
	queryErr := tx.QueryRowContext(ctx, QUERY_GET_PARTICIPANT_BY_NAME_AND_EVENT_ID, name, eventId).Scan(&e.EventId, &e.Name, &e.CreatedBy, &e.UpdatedBy, &e.CreatedAt, &e.UpdatedAt)
	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, queryErr
		}
	}

	return &e, tx.Commit()
}
