package data

import (
	"atta/usercenter/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type userCenterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.UserCenterRepo {
	return &userCenterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userCenterRepo) CreateGreeter(ctx context.Context, u *biz.UserCenter) error {
	return nil
}

func (r *userCenterRepo) UpdateGreeter(ctx context.Context, u *biz.UserCenter) error {
	return nil
}
