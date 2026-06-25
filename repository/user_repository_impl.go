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

// GetUser retrieves a user from database identified by their username, excluding password and deleted_at.
func (r *userRepository) GetUser(ctx context.Context, username string) (entity.User, error) {
	query := `
	SELECT uuid, username, fullname, email, photo, created_at, updated_at
	FROM users
	WHERE username = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	var user entity.User
	err := stmt.QueryRowContext(ctx, username).
		Scan(
			&user.UUID,
			&user.Username,
			&user.Fullname,
			&user.Email,
			&user.Photo,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	return user, err
}

// GetUserByUUID retrieves a user identified by uuid, excluding password and deleted_at.
func (r *userRepository) GetUserByUUID(ctx context.Context, uuid []byte) (entity.User, error) {
	query := `
	SELECT uuid, username, fullname, email, photo, created_at, updated_at
	FROM users
	WHERE uuid = ?
	LIMIT 1
	`
	stmt, _ := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	var uuidBinary []byte
	var user entity.User
	if err := stmt.QueryRowContext(ctx, uuid).
		Scan(
			&uuidBinary,
			&user.Username,
			&user.Fullname,
			&user.Email,
			&user.Photo,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
		return user, err
	}

	// uuid is stored in BINARY(16)
	// the struct field accepts string
	uuidStr, err := util.UUIDtoString(uuidBinary)
	if err != nil {
		return user, err
	}
	user.UUID = uuidStr

	return user, nil
}
