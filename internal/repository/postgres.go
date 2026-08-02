package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

// CREATE

func (r *PostgresRepository) Create(
	ctx context.Context,
	user model.User,
) (model.User, error) {

	query := `
	INSERT INTO users(name,email,age,password_hash)
	VALUES($1,$2,$3,$4)
	RETURNING id
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Age,
		user.PasswordHash,
	).Scan(&user.Id)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// UPDATE

func (r *PostgresRepository) Update(
	ctx context.Context,
	id int,
	user model.User,
) (model.User, error) {

	query := `
	UPDATE users
	SET
		name = $1,
		email = $2,
		age = $3,
		password_hash = $4
	WHERE id = $5
	RETURNING id
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Age,
		user.PasswordHash,
		id,
	).Scan(&user.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, ErrNotFound
	}

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// DELETE

func (r *PostgresRepository) Delete(
	ctx context.Context,
	id int,
) error {

	query := `
	DELETE FROM users
	WHERE id = $1
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

// GET BY ID

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id int,
) (model.User, error) {

	query := `
	SELECT
		id,
		name,
		email,
		age
	FROM users
	WHERE id = $1
	`

	var user model.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Age,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, ErrNotFound
	}

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// GET ALL

func (r *PostgresRepository) GetAll(
	ctx context.Context,
) ([]model.User, error) {

	query := `
	SELECT
		id,
		name,
		email,
		age
	FROM users
	ORDER BY id
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {

		var user model.User

		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Age,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
