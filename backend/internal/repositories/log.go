package repositories

import (
	"context"
	"database/sql"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LogRepository struct {
	db *sqlx.DB
}

// dbModelLogs is the struct that represents the log in the database.
type dbModelLogs struct {
	Id        sql.NullString `db:"id"`
	UserId    sql.NullString `db:"user_id"`
	Title     sql.NullString `db:"title"`
	Type      sql.NullString `db:"type"`
	Content   sql.NullString `db:"content"`
	CreatedAt sql.NullTime   `db:"created_at"`
}

// dbModelToAppModel converts dbModelLogs to domains.Log for application operations (e.g. return to client)
func (r *LogRepository) dbModelToAppModel(dbModel dbModelLogs) (log domains.Log) {
	log.Unmarshal(
		uuid.MustParse(dbModel.Id.String),
		uuid.MustParse(dbModel.UserId.String),
		dbModel.Title.String,
		dbModel.Type.String,
		dbModel.Content.String,
		dbModel.CreatedAt.Time,
	)
	return
}

// dbModelFromAppModel converts domains.Log to dbModelLogs for database operations (e.g. insert, update)
func (r *LogRepository) dbModelFromAppModel(domModel domains.Log) (dbModel dbModelLogs) {
	if domModel.ID() != uuid.Nil {
		dbModel.Id.String = domModel.ID().String()
		dbModel.Id.Valid = true
	}
	if domModel.UserID() != uuid.Nil {
		dbModel.UserId.String = domModel.UserID().String()
		dbModel.UserId.Valid = true
	}
	if domModel.Title() != "" {
		dbModel.Title.String = domModel.Title()
		dbModel.Title.Valid = true
	}
	if domModel.Type() != "" {
		dbModel.Type.String = domModel.Type()
		dbModel.Type.Valid = true
	}
	if domModel.Content() != "" {
		dbModel.Content.String = domModel.Content()
		dbModel.Content.Valid = true
	}
	if !domModel.CreatedAt().IsZero() {
		dbModel.CreatedAt.Time = domModel.CreatedAt()
		dbModel.CreatedAt.Valid = true
	}

	return
}

// dbModelFromAppModel converts domains.LogFilter to dbModelLogs for database operations (e.g. insert, update)
func (r *LogRepository) dbModelFromAppFilter(filter domains.LogFilter) (dbFilter dbModelLogs) {
	if filter.ID != uuid.Nil {
		dbFilter.Id.String = filter.ID.String()
		dbFilter.Id.Valid = true
	}
	if filter.UserID != uuid.Nil {
		dbFilter.UserId.String = filter.UserID.String()
		dbFilter.UserId.Valid = true
	}
	if filter.Title != "" {
		dbFilter.Title.String = filter.Title
		dbFilter.Title.Valid = true
	}
	if filter.LType != "" {
		dbFilter.Type.String = filter.LType
		dbFilter.Type.Valid = true
	}
	if filter.Content != "" {
		dbFilter.Content.String = filter.Content
		dbFilter.Content.Valid = true
	}

	return
}

func NewLogRepository(db *sqlx.DB) domains.ILogRepository {
	return &LogRepository{db: db}
}

// Devam
func (r *LogRepository) Filter(ctx context.Context, filter domains.LogFilter) (logs []domains.Log, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelLogs{}

	query := `
	SELECT
		*
	FROM t_logs
	WHERE
		(? IS NULL OR id = ?) AND
		(? IS NULL OR user_id = ?) AND
		(? IS NULL OR title LIKE CONCAT('%', ?, '%')) AND
		(? IS NULL OR type LIKE CONCAT('%', ?, '%')) AND
		(? IS NULL OR content LIKE CONCAT('%', ?, '%'))
	`

	err = r.db.Select(&dbResult, query, dbFilter.Id, dbFilter.Id, dbFilter.UserId, dbFilter.UserId, dbFilter.Title, dbFilter.Title, dbFilter.Type, dbFilter.Type, dbFilter.Content, dbFilter.Content)
	if err != nil {
		return
	}
	for _, dbModel := range dbResult {
		logs = append(logs, r.dbModelToAppModel(dbModel))
	}

	return
}

// Adds Log
func (r *LogRepository) Add(ctx context.Context, log *domains.Log) (err error) {
	// Checks the logs already in the db. If the log exists then we will not insert a new one.
	query := `
		SELECT
			EXISTS (
				SELECT 1
				FROM t_logs
				WHERE id = :id AND user_id = :user_id AND type = :type AND content = :content AND title = :title
			)
	`
	params := map[string]interface{}{
		"id":      log.ID,
		"user_id": log.UserID,
		"type":    log.Type,
		"content": log.Content,
		"title":   log.Title,
	}

	var exists bool
	err = r.db.GetContext(ctx, &exists, query, params)
	if err != nil {
		return err
	}
	if exists {
		return
	}

	dbModel := r.dbModelFromAppModel(*log)
	query = `
		INSERT INTO
			t_logs
		(id, user_id, title, type, content)
			VALUES
		(:id, :user_id, :title, :type, :content)
	`

	_, err = r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return
	}
	return
}

func (r *LogRepository) IsExists(ctx context.Context, log *domains.LogFilter) (exists bool, err error) {
	query := `
		SELECT
			EXISTS (
				SELECT 1
				FROM t_logs
				WHERE id = :id
			)
	`

	err = r.db.GetContext(ctx, &exists, query, log.ID)
	if err != nil {
		return false, err
	}
	return exists, nil
}
