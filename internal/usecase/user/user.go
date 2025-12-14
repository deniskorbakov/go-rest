package user

import "github.com/go-list-templ/grpc/internal/domain/repo"

type UseCase struct {
	repo repo.User
}

func New(repo repo.User) *UseCase {
	return &UseCase{repo: repo}
}
