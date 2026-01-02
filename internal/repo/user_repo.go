package repo

import (
	"context"
	"encoding/json"
	"github.com/go-list-templ/grpc/internal/domain/entity"
	"go.uber.org/zap"
	"time"

	"github.com/go-list-templ/grpc/pkg/redis"
)

const (
	CacheKeyAllUsers = "users:all"
)

type UserRepo struct {
	persistent UserPersistentRepo
	redis      redis.Redis
	ctx        context.Context
	logger     zap.Logger
}

func NewUserRepo(ctx context.Context, p UserPersistentRepo, r redis.Redis, logger zap.Logger) *UserRepo {
	return &UserRepo{persistent: p, redis: r, ctx: ctx, logger: logger}
}

func (u *UserRepo) All(ctx context.Context) ([]entity.User, error) {
	cachedData, err := u.redis.Get(ctx, CacheKeyAllUsers).Bytes()
	if err == nil {
		var users []entity.User
		if err = json.Unmarshal(cachedData, &users); err == nil {
			return users, nil
		}
	}

	users, err := u.persistent.All(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		data, err := json.Marshal(users)
		if err != nil {
			u.logger.Error("marshal users error", zap.Error(err))
		}

		err = u.redis.Set(ctx, CacheKeyAllUsers, data, 10*time.Minute).Err()
		if err != nil {
			u.logger.Error("redis set error", zap.Error(err))
		}
	}()

	return users, nil
}

func (u *UserRepo) Store(ctx context.Context, user entity.User) (entity.User, error) {
	err := u.persistent.Store(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	go func() {
		err = u.redis.Del(ctx, CacheKeyAllUsers).Err()
		if err != nil {
			u.logger.Error("redis del error", zap.Error(err))
		}
	}()

	return user, nil
}
