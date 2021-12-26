package repository

import (
	"context"
	"database/sql"

	"minibank/auth"
	"minibank/db/postgresql"
)

type UserRepository struct {
	q *postgresql.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{q: postgresql.New(db)}
}

func (r UserRepository) CreateUser(ctx context.Context, email string, password string) error {
	_, err := r.q.InsertUser(ctx, postgresql.InsertUserParams{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) GetUser(ctx context.Context, id int) (*auth.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r UserRepository) GetUserByEmail(ctx context.Context, email string) (*auth.User, error) {
	u, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	user := u.MapToEntity()

	return &user, nil
}
