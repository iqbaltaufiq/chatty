package repository

import (
	"context"
	"database/sql"

	"chatty/entity"
	"chatty/util"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) error {
	query := `
	INSERT INTO users (uuid, username, fullname, email, password, photo)
	VALUES (?,?,?,?,?,?)
	`

	// column `uuid` has BINARY(16) data type.
	// convert UUID string to binary before storing.
	binaryID, err := util.UUIDtoBinary(user.UUID)
	if err != nil {
		return err
	}

	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		binaryID,
		user.Username,
		user.Fullname,
		user.Email,
		user.Password,
		user.Photo,
	)
	return err
}

// GetPassword should only be used for authentication purpose.
// It returns the user's password.
func (r *userRepository) GetPassword(ctx context.Context, username string) (string, error) {
	query := `
	SELECT password
	FROM users
	WHERE username = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	var password string
	row := stmt.QueryRowContext(ctx, username)
	err := row.Scan(&password)
	return password, err
}
