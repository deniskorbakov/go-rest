package persistent

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
	"github.com/go-list-templ/grpc/internal/domain/vo"
	"github.com/go-list-templ/grpc/internal/infra/persistent/postgres"
)

type UserPostgresRepo struct {
	*postgres.Postgres
}

func NewUserPostgresRepo(postgres *postgres.Postgres) *UserPostgresRepo {
	return &UserPostgresRepo{postgres}
}

func (r *UserPostgresRepo) Store(ctx context.Context, user entity.User) error {
	query := `
		INSERT INTO users (id, name, email, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.Pool.Exec(ctx, query,
		user.ID,
		user.Username,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgresRepo) Change(ctx context.Context, user entity.User) (entity.User, error) {
	var updatedUser entity.User

	query := `
		UPDATE users SET name = $2, email = $3, updated_at = $4 WHERE id = $1	
	`

	err := r.Pool.QueryRow(ctx, query,
		user.ID,
		user.Username,
		user.Email,
		user.UpdatedAt,
	).Scan(&updatedUser)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (r *UserPostgresRepo) Destroy(ctx context.Context, id vo.ID) error {
	_, err := r.Pool.Exec(ctx, "DELETE FROM users WHERE id = $1", id)

	return err
}

func (r *UserPostgresRepo) GetById(ctx context.Context, id vo.ID) (entity.User, error) {
	var user entity.User

	err := r.Pool.QueryRow(ctx, `SELECT * FROM users WHERE id = $1`, id).Scan(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserPostgresRepo) All(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	err := r.Pool.QueryRow(ctx, "SELECT * FROM users").Scan(&users)
	if err != nil {
		return users, err
	}

	return users, nil
}
