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
	QUERY_CREATE_COMMENT                   string = "INSERT INTO present_comment (presentId, content, createdBy, updatedBy, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)"
	QUERY_GET_COMMENTS_BY_PRESENT_ID       string = "SELECT * FROM present_comment as c WHERE c.presentId = ?"
	QUERY_GET_COMMENT_COUNT_BY_PRESENT_IDS string = "SELECT DISTINCT v.presentId, COUNT(v.presentId) AS commentCount FROM present_comment AS v WHERE v.presentId IN (%s) GROUP BY v.presentId"
)

type CommentRepository struct {
	dbConn *database.Connection
}

func NewCommentRepository(dbConn *database.Connection) *CommentRepository {
	return &CommentRepository{dbConn: dbConn}
}

func (repository *CommentRepository) CreateComment(comment models.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()

	return repository.dbConn.ExecuteInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, execError := tx.ExecContext(ctx, QUERY_CREATE_COMMENT, comment.PresentId, comment.Content, comment.CreatedBy, comment.UpdatedBy, now, now)
		if execError != nil {
			return execError
		}

		return nil
	})
}

func (repository *CommentRepository) GetCommentsByPresentId(presentId uuid.UUID) ([]*models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	rows, queryErr := tx.QueryContext(ctx, QUERY_GET_COMMENTS_BY_PRESENT_ID, presentId)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var c models.Comment
		err := rows.Scan(&c.PresentId, &c.Content, &c.CreatedBy, &c.UpdatedBy, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &c)
	}

	return comments, tx.Commit()
}

func (repository *CommentRepository) GetCommentCountMapByPresentIds(presentIds []string) (map[string]int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, txErr := repository.dbConn.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return nil, txErr
	}

	defer tx.Rollback()

	placeholders := strings.Join(strings.Split(strings.Repeat("?", len(presentIds)), ""), ",")
	query := fmt.Sprintf(QUERY_GET_COMMENT_COUNT_BY_PRESENT_IDS, placeholders)

	args := make([]interface{}, len(presentIds)+1)
	for i, id := range presentIds {
		args[i] = id
	}

	rows, queryErr := tx.QueryContext(ctx, query, args...)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	presentIdsToCommentCount := make(map[string]int32)
	for rows.Next() {
		var v struct {
			PresentId    string
			CommentCount int32
		}
		err := rows.Scan(&v.PresentId, &v.CommentCount)
		if err != nil {
			return nil, err
		}
		presentIdsToCommentCount[v.PresentId] = v.CommentCount
	}

	return presentIdsToCommentCount, tx.Commit()

}
