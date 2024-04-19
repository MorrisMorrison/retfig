package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/MorrisMorrison/retfig/database"
	"github.com/MorrisMorrison/retfig/models"
	uuid "github.com/satori/go.uuid"
)

const (
	QUERY_CREATE_VOTE                                 string = "INSERT INTO present_vote (presentId, type, createdBy, updatedBy, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)"
	QUERY_GET_VOTE_BY_PRESENT_ID_AND_USERNAME         string = "SELECT * FROM present_vote AS V WHERE presentId = ? and CreatedBy = ?"
	QUERY_DELETE_VOTE_BY_PRESENT_ID_AND_USERNAME      string = "DELETE FROM present_vote AS V WHERE presentId = ? and CreatedBy = ?"
	QUERY_GET_VOTE_COUNT_BY_PRESENT_IDS_AND_VOTE_TYPE string = "SELECT DISTINCT v.presentId, COUNT(v.presentId) AS voteCount FROM present_vote AS v WHERE v.presentId IN (%s) AND v.type = ? GROUP BY v.presentId"
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

func (repository *VoteRepository) GetVoteByPresentIdAndUser(presentId uuid.UUID, username string) (*models.Vote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var v models.Vote
	queryErr := tx.QueryRowContext(ctx, QUERY_GET_VOTE_BY_PRESENT_ID_AND_USERNAME, presentId, username).Scan(&v.PresentId, &v.Type, &v.CreatedBy, &v.UpdatedBy, &v.CreatedAt, &v.UpdatedAt)
	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, queryErr
		}
	}

	return &v, tx.Commit()
}

func (repository *VoteRepository) DeleteVoteByPresentIdAndUsername(presentId uuid.UUID, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repository.dbConn.ExecuteInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, execError := tx.ExecContext(ctx, QUERY_DELETE_VOTE_BY_PRESENT_ID_AND_USERNAME, presentId, username)
		if execError != nil {
			return execError
		}

		return nil
	})
}

func (repository *VoteRepository) GetVoteCountMapByPresentIdsAndVoteType(presentIds []string, voteType models.VoteType) (map[string]int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	placeholders := strings.Join(strings.Split(strings.Repeat("?", len(presentIds)), ""), ",")
	query := fmt.Sprintf(QUERY_GET_VOTE_COUNT_BY_PRESENT_IDS_AND_VOTE_TYPE, placeholders)

	args := make([]interface{}, len(presentIds)+1)
	for i, id := range presentIds {
		args[i] = id
	}
	args[len(presentIds)] = string(voteType)

	rows, queryErr := tx.QueryContext(ctx, query, args...)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	presentIdsToVoteCount := make(map[string]int32)
	for rows.Next() {
		var v struct {
			PresentId string
			VoteCount int32
		}
		err := rows.Scan(&v.PresentId, &v.VoteCount)
		if err != nil {
			return nil, err
		}
		presentIdsToVoteCount[v.PresentId] = v.VoteCount
	}

	return presentIdsToVoteCount, tx.Commit()
}
