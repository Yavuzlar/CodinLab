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
	ID            sql.NullString `db:"id"`
	UserID        sql.NullString `db:"user_id"`
	ProgrammingID sql.NullInt32  `db:"programming_id"`
	LabPathID     sql.NullInt32  `db:"lab_path_id"`
	Type          sql.NullString `db:"type"`
	Content       sql.NullString `db:"content"`
	CreatedAt     sql.NullTime   `db:"created_at"`
}

// lab and road numbers solved day by day
type dbModelSolutionsByDay struct {
	Date  string `db:"date"`
	Count int    `db:"count"`
}

// SolutionsByLanguage represents the total count from lab and road solutions for each language.
type dbModelSolutionsByProgrammingLanguage struct {
	ProgrammingID  int32 `db:"programming_id"`
	TotalLabCount  int   `db:"total_lab_count"`
	TotalRoadCount int   `db:"total_road_count"`
	TotalCount     int   `db:"total_count"`
}

func (r *LogRepository) dbModelSolutionsByDayToAppModel(dbModelSolutionsByDay dbModelSolutionsByDay) (solutionsByDay domains.SolutionsByDay) {
	date, _ := time.Parse("2006-01-02", dbModelSolutionsByDay.Date)

	solutionsByDay.SetDate(date)
	solutionsByDay.SetCount(dbModelSolutionsByDay.Count)
	solutionsByDay.SetLevel(0)
	return
}

func (r *LogRepository) dbModelSolutionsCountToAppModel(dbModel dbModelSolutionsByProgrammingLanguage) (appModel domains.SolutionsByProgramming) {
	appModel.SetLabCount(dbModel.TotalLabCount)
	appModel.SetProgrammingID(dbModel.ProgrammingID)
	appModel.SetRoadCount(dbModel.TotalRoadCount)
	appModel.SetTotalCount(dbModel.TotalCount)
	return
}

// dbModelToAppModel converts dbModelLogs to domains.Log for application operations (e.g. return to client)
func (r *LogRepository) dbModelToAppModel(dbModel dbModelLogs) (log domains.Log) {
	log.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.UserID.String),
		dbModel.Type.String,
		dbModel.Content.String,
		dbModel.ProgrammingID.Int32,
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
	if domModel.ProgrammingID() != 0 {
		dbModel.ProgrammingID.Int32 = domModel.ProgrammingID()
		dbModel.ProgrammingID.Valid = true
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
	if filter.ProgrammingID != 0 {
		dbFilter.ProgrammingID.Int32 = filter.ProgrammingID
		dbFilter.ProgrammingID.Valid = true
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
		(? IS NULL OR programming_id = ?) AND
		(? IS NULL OR lab_path_id = ?) AND
		(? IS NULL OR type LIKE CONCAT('%', ?, '%')) AND
		(? IS NULL OR content LIKE CONCAT('%', ?, '%'))
	`

	err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.ID, dbFilter.UserID, dbFilter.UserID, dbFilter.ProgrammingID, dbFilter.ProgrammingID, dbFilter.LabPathID, dbFilter.LabPathID, dbFilter.Type, dbFilter.Type, dbFilter.Content, dbFilter.Content)
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
	ok, err := r.IsExists(ctx, log)
	if err != nil {
		return
	}
	if ok {
		return
	}

	dbModel := r.dbModelFromAppModel(*log)
	query := `
		INSERT INTO
			t_logs
		(id, user_id, programming_id, lab_path_id, type, content)
			VALUES
		(:id, :user_id, :programming_id, :lab_path_id, :type, :content)
	`

	_, err = r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return
	}
	return
}

func (r *LogRepository) IsExists(ctx context.Context, log *domains.Log) (exists bool, err error) {
	// if log type is user you can add multiple log.
	if canMultipleLogExists(log.Type()) {
		return false, nil
	}

	// Checks the logs already in the db. If the log exists then we will not insert a new one.
	query := `
		SELECT
			EXISTS (
				SELECT 1
				FROM t_logs
				WHERE 
					user_id = :user_id AND 
					((programming_id IS NULL AND :programming_id IS NULL) OR (programming_id = :programming_id)) AND
					type = :type AND 
					content = :content AND 
					((lab_path_id IS NULL AND :lab_path_id IS NULL) OR (lab_path_id = :lab_path_id))
			)
	`

	params := r.dbModelFromAppModel(*log)

	err = r.db.GetContext(ctx, &exists, query, params.UserID, params.ProgrammingID, params.Type, params.Content, params.LabPathID)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}

	return exists, nil
}

// CountSolutionsByDay counts the number of lab and road solutions completed each day.
func (r *LogRepository) CountSolutionsByDay(ctx context.Context, year string) (solutions []domains.SolutionsByDay, err error) {
	var dbModelSolutions []dbModelSolutionsByDay

	query := `
	SELECT 
		DATE(created_at) AS date,
		SUM(CASE WHEN type = 'Lab' AND content = 'Completed' THEN 1 ELSE 0 END) AS count
	FROM 
		t_logs
	WHERE
		content = 'Completed'
		AND strftime('%Y', created_at) = ?
	GROUP BY 
		DATE(created_at)
	ORDER BY 
		DATE(created_at) DESC
	`

	err = r.db.SelectContext(ctx, &dbModelSolutions, query, year)
	if err != nil {
		return nil, err
	}

	for _, dbModelSolution := range dbModelSolutions {
		solutions = append(solutions, r.dbModelSolutionsByDayToAppModel(dbModelSolution))
	}

	return solutions, nil
}

// CountSolutionsByLanguageLast7Days counts the solved lab and road solutions in the last 7 days for each language.
func (r *LogRepository) CountSolutionsByProgrammingLast7Days(ctx context.Context) (solutionsCount []domains.SolutionsByProgramming, err error) {
	query := `
    SELECT
        l1.programming_id,
        COUNT(CASE WHEN l1.type = 'Lab' THEN 1 ELSE NULL END) AS total_lab_count,
        COUNT(CASE WHEN l1.type = 'Path' THEN 1 ELSE NULL END) AS total_road_count,
		COUNT(CASE WHEN l1.type IN ('Lab', 'Path') THEN 1 ELSE NULL END) AS total_count
    FROM
        t_logs l1
    JOIN
        t_logs l2 ON l1.user_id = l2.user_id
                 AND l1.programming_id = l2.programming_id
                 AND l1.lab_path_id = l2.lab_path_id
                 AND l1.type = l2.type
                 AND l1.content = 'Started'
                 AND l2.content = 'Completed'
    WHERE
        l2.type IN ('Path', 'Lab')
        AND l2.created_at >= DATE('now', '-7 days')
    GROUP BY
        l1.programming_id
`

	var dbModelSolutionsCount []dbModelSolutionsByProgrammingLanguage

	err = r.db.SelectContext(ctx, &dbModelSolutionsCount, query)
	if err != nil {
		return nil, err
	}

	for _, result := range dbModelSolutionsCount {
		solutionsCount = append(solutionsCount, r.dbModelSolutionsCountToAppModel(result))
	}

	return solutionsCount, nil
}

func canMultipleLogExists(logType string) bool {
	return logType == domains.TypeUser
}
