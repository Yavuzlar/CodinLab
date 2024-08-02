package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LogRepository struct {
	db *sqlx.DB
}

// dbModelLogs is the struct that represents the log in the database.
type dbModelLogs struct {
	ID         sql.NullString `db:"id"`
	UserID     sql.NullString `db:"user_id"`
	LanguageID sql.NullInt32  `db:"language_id"`
	LabPathID  sql.NullInt32  `db:"lab_path_id"`
	Type       sql.NullString `db:"type"`
	Content    sql.NullString `db:"content"`
	CreatedAt  sql.NullTime   `db:"created_at"`
}

// lab and road numbers solved day by day
// author: yasir
type dbModelSolutionsByDay struct {
	Date      string `db:"date"`
	RoadCount int    `db:"road_count"`
	LabCount  int    `db:"lab_count"`
}

// author: yasir
func (r *LogRepository) dbModelSolutionsByDayToAppModel(dbModelSolutionsByDay dbModelSolutionsByDay) (solutionsByDay domains.SolutionsByDay, err error) {
	date, parseErr := time.Parse("2006-01-02", dbModelSolutionsByDay.Date)
	if parseErr != nil {
		return domains.SolutionsByDay{}, parseErr
	}
	return domains.SolutionsByDay{
		Date:      date,
		RoadCount: dbModelSolutionsByDay.RoadCount,
		LabCount:  dbModelSolutionsByDay.LabCount,
	}, nil
}

// dbModelToAppModel converts dbModelLogs to domains.Log for application operations (e.g. return to client)
func (r *LogRepository) dbModelToAppModel(dbModel dbModelLogs) (log domains.Log) {
	log.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.UserID.String),
		dbModel.Type.String,
		dbModel.Content.String,
		dbModel.LanguageID.Int32,
		dbModel.LabPathID.Int32,
		dbModel.CreatedAt.Time,
	)
	return
}

// dbModelFromAppModel converts domains.Log to dbModelLogs for database operations (e.g. insert, update)
func (r *LogRepository) dbModelFromAppModel(domModel domains.Log) (dbModel dbModelLogs) {
	if domModel.ID() != uuid.Nil {
		dbModel.ID.String = domModel.ID().String()
		dbModel.ID.Valid = true
	}
	if domModel.UserID() != uuid.Nil {
		dbModel.UserID.String = domModel.UserID().String()
		dbModel.UserID.Valid = true
	}
	if domModel.LanguageID() != 0 {
		dbModel.LanguageID.Int32 = domModel.LanguageID()
		dbModel.LanguageID.Valid = true
	}
	if domModel.LabPathID() != 0 {
		dbModel.LabPathID.Int32 = domModel.LabPathID()
		dbModel.LabPathID.Valid = true
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
		dbFilter.ID.String = filter.ID.String()
		dbFilter.ID.Valid = true
	}
	if filter.UserID != uuid.Nil {
		dbFilter.UserID.String = filter.UserID.String()
		dbFilter.UserID.Valid = true
	}
	if filter.LanguageID != 0 {
		dbFilter.LanguageID.Int32 = filter.LanguageID
		dbFilter.LanguageID.Valid = true
	}
	if filter.LabPathID != 0 {
		dbFilter.LabPathID.Int32 = filter.LabPathID
		dbFilter.LabPathID.Valid = true
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
		(? IS NULL OR language_id = ?) AND
		(? IS NULL OR lab_path_id = ?) AND
		(? IS NULL OR type LIKE CONCAT('%', ?, '%')) AND
		(? IS NULL OR content LIKE CONCAT('%', ?, '%'))
	`

	err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.ID, dbFilter.UserID, dbFilter.UserID, dbFilter.LanguageID, dbFilter.LanguageID, dbFilter.LabPathID, dbFilter.LabPathID, dbFilter.Type, dbFilter.Type, dbFilter.Content, dbFilter.Content)
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
            WHERE 
                user_id = :user_id AND 
                ((language_id IS NULL AND :language_id IS NULL) OR (language_id = :language_id)) AND
                type = :type AND 
                content = :content AND 
                ((lab_path_id IS NULL AND :lab_path_id IS NULL) OR (lab_path_id = :lab_path_id))
        )
`

	params := r.dbModelFromAppModel(*log)

	var exists bool
	err = r.db.GetContext(ctx, &exists, query, params.UserID, params.LanguageID, params.Type, params.Content, params.LabPathID)
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
		(id, user_id, language_id, lab_path_id, type, content)
			VALUES
		(:id, :user_id, :language_id, :lab_path_id, :type, :content)
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

// CountSolutionsByDay counts the number of lab and road solutions completed each day.
func (r *LogRepository) CountSolutionsByDay(ctx context.Context) (solutions []domains.SolutionsByDay, err error) {
	var dbModelSolutions []dbModelSolutionsByDay

	query := `
    SELECT 
        DATE(created_at) AS date,
        SUM(CASE WHEN type = 'Road' AND content = 'Completed' THEN 1 ELSE 0 END) AS road_count,
        SUM(CASE WHEN type = 'Lab' AND content = 'Completed' THEN 1 ELSE 0 END) AS lab_count
    FROM 
        t_logs
    WHERE
        content = 'Completed'
    GROUP BY 
        DATE(created_at)
    ORDER BY 
        DATE(created_at) DESC
    `

	err = r.db.SelectContext(ctx, &dbModelSolutions, query)
	if err != nil {
		return nil, err
	}

	for _, dbModelSolution := range dbModelSolutions {
		appModelSolution, parseErr := r.dbModelSolutionsByDayToAppModel(dbModelSolution)
		if parseErr != nil {
			return nil, parseErr
		}
		solutions = append(solutions, appModelSolution)
	}

	return solutions, nil
}

// SolutionsHoursByLanguage represents the total hours spent on lab and road solutions for each language.
type dbModelSolutionsHoursByLanguage struct {
	LanguageID     int32   `db:"language_id"`
	TotalLabHours  float64 `db:"total_lab_hours"`
	TotalRoadHours float64 `db:"total_road_hours"`
}

// CountSolutionsHoursByLanguageLast7Days counts the total hours spent on lab and road solutions in the last 7 days for each language.
func (r *LogRepository) CountSolutionsHoursByLanguageLast7Days(ctx context.Context) (solutionsHours []domains.SolutionsHoursByLanguage, err error) {
	query := `
	SELECT
		l1.language_id,
		SUM(CASE WHEN l1.type = 'Lab' THEN (JULIANDAY(l2.created_at) - JULIANDAY(l1.created_at)) * 24 ELSE 0 END) AS total_lab_hours,
		SUM(CASE WHEN l1.type = 'Road' THEN (JULIANDAY(l2.created_at) - JULIANDAY(l1.created_at)) * 24 ELSE 0 END) AS total_road_hours
	FROM
		t_logs l1
	JOIN
		t_logs l2 ON l1.user_id = l2.user_id
	             AND l1.language_id = l2.language_id
	             AND l1.lab_path_id = l2.lab_path_id
	             AND l1.type = l2.type
	             AND l1.content = 'Started'
	             AND l2.content = 'Completed'
	             AND l1.created_at < l2.created_at
	WHERE
		l1.type IN ('Road', 'Lab')
		AND l1.created_at >= DATE('now', '-7 days')
	GROUP BY
		l1.language_id
	`

	var dbModelSolutionsHours []dbModelSolutionsHoursByLanguage

	err = r.db.SelectContext(ctx, &dbModelSolutionsHours, query)
	if err != nil {
		return nil, err
	}

	for _, result := range dbModelSolutionsHours {
		solutionsHours = append(solutionsHours, domains.SolutionsHoursByLanguage{
			LanguageID: result.LanguageID,
			LabHours:   result.TotalLabHours,
			RoadHours:  result.TotalRoadHours,
		})
	}

	return solutionsHours, nil
}
