package repositories

import (
	"context"
	"database/sql"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

// dbModelUsers is the struct that represents the user in the database.
type dbModelUsers struct {
	Id            sql.NullString `db:"id"`
	Username      sql.NullString `db:"username"`
	Password      sql.NullString `db:"password"`
	Name          sql.NullString `db:"name"`
	Surname       sql.NullString `db:"surname"`
	Role          sql.NullString `db:"role"`
	GithubProfile sql.NullString `db:"github_profile"`
	TotalPoints   sql.NullInt32  `db:"total_points"`
	CreatedAt     sql.NullTime   `db:"created_at"`
}

// dbModelToAppModel converts dbModelUsers to domains.User for application operations (e.g. return to client)
func (r *UserRepository) dbModelToAppModel(dbModel dbModelUsers) (user domains.User) {
	user.Unmarshal(
		uuid.MustParse(dbModel.Id.String),
		dbModel.Username.String,
		dbModel.Password.String,
		dbModel.Name.String,
		dbModel.Surname.String,
		dbModel.Role.String,
		dbModel.GithubProfile.String,
		dbModel.TotalPoints.Int32,
		dbModel.CreatedAt.Time,
	)
	return
}

// dbModelFromAppModel converts domains.User to dbModelUsers for database operations (e.g. insert, update)
func (r *UserRepository) dbModelFromAppModel(domModel domains.User) (dbModel dbModelUsers) {
	if domModel.ID() != uuid.Nil {
		dbModel.Id.String = domModel.ID().String()
		dbModel.Id.Valid = true
	}
	if domModel.Username() != "" {
		dbModel.Username.String = domModel.Username()
		dbModel.Username.Valid = true
	}
	if domModel.Password() != "" {
		dbModel.Password.String = domModel.Password()
		dbModel.Password.Valid = true
	}
	if domModel.Name() != "" {
		dbModel.Name.String = domModel.Name()
		dbModel.Name.Valid = true
	}
	if domModel.Surname() != "" {
		dbModel.Surname.String = domModel.Surname()
		dbModel.Surname.Valid = true
	}
	if domModel.Role() != "" {
		dbModel.Role.String = domModel.Role()
		dbModel.Role.Valid = true
	}
	if domModel.GithubProfile() != "" {
		dbModel.GithubProfile.String = domModel.GithubProfile()
		dbModel.GithubProfile.Valid = true
	}
	if domModel.TotalPoints() != 0 {
		dbModel.TotalPoints.Int32 = domModel.TotalPoints()
		dbModel.TotalPoints.Valid = true
	}
	if !domModel.CreatedAt().IsZero() {
		dbModel.CreatedAt.Time = domModel.CreatedAt()
		dbModel.CreatedAt.Valid = true
	}
	return
}

// dbModelFromAppFilter converts domains.UserFilter to dbModelUsers for database operations (e.g. select)
func (r *UserRepository) dbModelFromAppFilter(filter domains.UserFilter) (dbFilter dbModelUsers) {
	if filter.Id != uuid.Nil {
		dbFilter.Id.String = filter.Id.String()
		dbFilter.Id.Valid = true
	}
	if filter.Username != "" {
		dbFilter.Username.String = filter.Username
		dbFilter.Username.Valid = true
	}
	if filter.Name != "" {
		dbFilter.Name.String = filter.Name
		dbFilter.Name.Valid = true
	}
	if filter.Surname != "" {
		dbFilter.Surname.String = filter.Surname
		dbFilter.Surname.Valid = true
	}
	if filter.Role != "" {
		dbFilter.Role.String = filter.Role
		dbFilter.Role.Valid = true
	}
	return
}

func NewUserRepository(db *sqlx.DB) domains.IUserRepository {
	return &UserRepository{db: db}
}

// Filter returns users that match the filter
func (r *UserRepository) Filter(ctx context.Context, filter domains.UserFilter, limit, page int64) (users []domains.User, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelUsers{}
	query := `
		SELECT
			*
		FROM t_users
		WHERE
		  (? IS NULL OR id = ?) AND
		  (? IS NULL OR username like ('%' || ? || '%')) AND
		  (? IS NULL OR name like ('%' || ? || '%')) AND
		  (? IS NULL OR surname like ('%' || ? || '%')) AND
		  (? IS NULL OR role = ?) 
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`
	err = r.db.Select(&dbResult, query, dbFilter.Id, dbFilter.Id, dbFilter.Username, dbFilter.Username, dbFilter.Name, dbFilter.Name, dbFilter.Surname, dbFilter.Surname, dbFilter.Role, dbFilter.Role, limit, (page-1)*limit)
	if err != nil {
		return
	}
	for _, dbModel := range dbResult {
		users = append(users, r.dbModelToAppModel(dbModel))
	}
	return
}

func (r *UserRepository) Add(ctx context.Context, user *domains.User) (err error) {
	// Converting User model to dbModel
	dbModel := r.dbModelFromAppModel(*user)
	query := `
		INSERT INTO 
			t_users
		(id, username, password, name, surname, role, github_profile)
			VALUES
		(:id, :username, :password, :name, :surname, :role, :github_profile)
	`

	_, err = r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return
	}
	return
}

func (r *UserRepository) Update(ctx context.Context, user *domains.User) (err error) {
	dbModel := r.dbModelFromAppModel(*user)
	query := `
		UPDATE
         t_users
		SET
			username = COALESCE(:username, username),
			password = COALESCE(:password, password),
			github_profile = COALESCE(:github_profile, github_profile),
			total_points = COALESCE(:total_points, total_points)
		WHERE
			id = :id

	`
	_, err = r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return
	}
	return
}

func (r *UserRepository) Delete(ctx context.Context, userID uuid.UUID) (err error) {
	query := `
		DELETE FROM
			t_users
		WHERE
			id = ?
	`
	_, err = r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return
	}
	return
}
