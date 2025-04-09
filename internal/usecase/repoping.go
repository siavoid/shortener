package usecase

import (
	"context"
	"time"
)

func (u *UseCase) RepoPing(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return u.urlStore.Ping(ctx)
}
