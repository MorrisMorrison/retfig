package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/MorrisMorrison/retfig/persistence/database"
	"github.com/MorrisMorrison/retfig/persistence/models"
	uuid "github.com/satori/go.uuid"
)

const (
	QUERY_CREATE_CLAIM               string = "INSERT INTO present_claim (presentId, createdBy, updatedBy, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?)"
	QUERY_GET_CLAIM_BY_PRESENT_ID    string = "SELECT * FROM present_claim AS c WHERE c.presentId = ?"
	QUERY_GET_CLAIMS_BY_PRESENT_IDS  string = "SELECT * FROM present_claim AS c WHERE c.presentId IN (%s)"
	QUERY_DELETE_CLAIM_BY_PRESENT_ID string = "DELETE FROM present_claim AS c WHERE c.presentId = ?"
)

type ClaimRepository struct {
	dbConn *database.Connection
}

func NewClaimRepository(dbConn *database.Connection) *ClaimRepository {
	return &ClaimRepository{dbConn: dbConn}
}

func (repository *ClaimRepository) CreateClaim(claim models.Claim) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()

	return repository.dbConn.ExecuteInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, execError := tx.ExecContext(ctx, QUERY_CREATE_CLAIM, claim.PresentId, claim.CreatedBy, claim.UpdatedBy, now, now)
		if execError != nil {
			return execError
		}

		return nil
	})
}

func (repository *ClaimRepository) GetClaimByPresentId(presentId uuid.UUID) (*models.Claim, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	var c models.Claim
	queryErr := tx.QueryRowContext(ctx, QUERY_GET_CLAIM_BY_PRESENT_ID, presentId).Scan(&c.PresentId, &c.CreatedBy, &c.UpdatedBy, &c.CreatedAt, &c.UpdatedAt)
	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, queryErr
		}
	}

	return &c, tx.Commit()
}

func (repository *ClaimRepository) DeleteClaimByPresentId(presentId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repository.dbConn.ExecuteInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, execError := tx.ExecContext(ctx, QUERY_DELETE_CLAIM_BY_PRESENT_ID, presentId)
		if execError != nil {
			return execError
		}

		return nil
	})
}

func (repository *ClaimRepository) GetClaimsByPresentIds(presentIds []string) (map[string]*models.Claim, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	placeholders := strings.Join(strings.Split(strings.Repeat("?", len(presentIds)), ""), ",")
	query := fmt.Sprintf(QUERY_GET_CLAIMS_BY_PRESENT_IDS, placeholders)

	args := make([]interface{}, len(presentIds))
	for i, id := range presentIds {
		args[i] = id
	}
	rows, queryErr := tx.QueryContext(ctx, query, args...)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	presentIdsToClaims := make(map[string]*models.Claim)
	for rows.Next() {
		var c models.Claim
		err := rows.Scan(&c.PresentId, &c.CreatedBy, &c.UpdatedBy, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		presentIdsToClaims[c.PresentId.String()] = &c
	}

	return presentIdsToClaims, tx.Commit()
}
