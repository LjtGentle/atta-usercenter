package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserCenter struct {}

type UserCenterRepo interface {
	CreateGreeter(context.Context, *UserCenter) error
	UpdateGreeter(context.Context, *UserCenter) error
}

type UserCenterCase struct {
	repo UserCenterRepo
	log  *log.Helper
}

func NewUserCenterUseCase(repo UserCenterRepo, logger log.Logger) *UserCenterCase {
	return &UserCenterCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserCenterCase) Create(ctx context.Context, u *UserCenter) error {
	return uc.repo.CreateGreeter(ctx, u)
}

func (uc *UserCenterCase) Update(ctx context.Context, u *UserCenter) error {
	return uc.repo.UpdateGreeter(ctx, u)
}
