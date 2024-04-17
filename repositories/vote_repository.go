package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/MorrisMorrison/retfig/database"
	"github.com/MorrisMorrison/retfig/models"
)

const (
	QUERY_CREATE_VOTE string = "INSERT INTO present_vote (presentId, type, createdBy, updatedBy, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)"
)

type VoteRepository struct {
	dbConn *database.Connection
}

func NewVoteRepository(dbConn *database.Connection) *VoteRepository {
	return &VoteRepository{dbConn: dbConn}
}

func (repository *VoteRepository) CreateVote(vote models.Vote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()

	return repository.dbConn.ExecuteInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, execError := tx.ExecContext(ctx, QUERY_CREATE_VOTE, vote.PresentId, vote.Type, vote.CreatedBy, vote.UpdatedBy, now, now)
		if execError != nil {
			return execError
		}

		return nil
	})
}
