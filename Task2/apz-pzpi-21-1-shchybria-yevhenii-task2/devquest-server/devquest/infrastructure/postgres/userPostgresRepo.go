package postgres

import (
	"context"
	"database/sql"
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"devquest-server/devquest/infrastructure"

	"github.com/google/uuid"
)

type userPostgresRepo struct {
	db infrastructure.Database
}

func NewUserPostgresRepo(db infrastructure.Database) repositories.UserRepo {
	return &userPostgresRepo{db: db}
}

func (u *userPostgresRepo) GetUserByUsername(username string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, first_name, last_name, username, password_hash, role_id, company_id
		FROM users
		WHERE username = $1
	`

	row := u.db.GetDB().QueryRowContext(ctx, query, username)

	var user entities.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.PasswordHash, &user.RoleID, &user.CompanyID, &user.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userPostgresRepo) GetRoleByID(roleID uuid.UUID) (*entities.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, title
		FROM roles
		WHERE id = $1
	`

	row := u.db.GetDB().QueryRowContext(ctx, query, roleID)
	
	var role entities.Role
	err := row.Scan(&role.ID, &role.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &role, nil
}

func (u *userPostgresRepo) InsertUser(user *models.InsertUserDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.db.GetDBTimeout())
	defer cancel()

	execute := `
		INSERT INTO users
		(id, username, first_name, last_name, password_hash, role_id, company_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := u.db.GetDB().ExecContext(ctx, execute, user.ID, user.Username, user.FirstName, user.LastName, user.PasswordHash, user.RoleID, user.CompanyID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userPostgresRepo) CheckUserRole(userID uuid.UUID, roleTitle string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT u.id
		FROM users u
		WHERE u.id = $1
		AND u.role_id IN (SELECT id FROM roles WHERE title = $2)
	`

	row := u.db.GetDB().QueryRowContext(ctx, query, userID, roleTitle)

	var existingUserID uuid.UUID
	err := row.Scan(&existingUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
