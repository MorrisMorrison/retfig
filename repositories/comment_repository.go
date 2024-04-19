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
	QUERY_CREATE_COMMENT             string = "INSERT INTO present_comment (presentId, content, createdBy, updatedBy, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)"
	QUERY_GET_COMMENTS_BY_PRESENT_ID string = "SELECT * FROM present_comment as c WHERE c.presentId = ?"
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

func (repository *CommentRepository) GetCommentsByPresentId(presentId uuid.UUID) ([]models.Comment, error) {
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

	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		err := rows.Scan(&c.PresentId, &c.Content, &c.CreatedBy, &c.UpdatedBy, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}

	return comments, tx.Commit()
}
