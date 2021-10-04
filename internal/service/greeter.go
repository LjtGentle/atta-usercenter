package service

import (
	"context"

	v1 "atta/usercenter/api/usercenter/v1"
	"atta/usercenter/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type UserCenterService struct {
	v1.UnimplementedLoginServer
	v1.UnimplementedApplyServer
	v1.UnimplementedRegisterServer
	uc *biz.UserCenterCase
	log *log.Helper
}

// NewUserCenterService new a UserCenter service.
func NewUserCenterService(uc *biz.UserCenterCase, logger log.Logger) *UserCenterService {
	return &UserCenterService{uc: uc, log: log.NewHelper(logger)}
}


func (s *UserCenterService) RealNameVerification (ctx context.Context, in *v1.RealNameVerificationRequest) ( *v1.RealNameVerificationReply, error) {
	//TODO to be completed
	return nil,nil
}

func (s *UserCenterService) BecomeGameCoach (ctx context.Context, in *v1.BecomeGameCoachRequest) (*v1.BecomeGameCoachReply,error) {
	//TODO to be completed
	return nil,nil
}

func (s *UserCenterService) OrdinaryLogin (ctx context.Context, in *v1.OrdinaryLoginRequest ) (*v1.OrdinaryLoginReply,error) {
	//TODO to be completed
	return nil,nil
}

func (s *UserCenterService) AdminLogin (ctx context.Context, in *v1.AdminLoginRequest) (*v1.AdminLoginReply,error) {
	//TODO to be completed
	return nil,nil
}

func (s *UserCenterService) OrdinaryRegister (ctx context.Context, in *v1.OrdinaryRegisterRequest) (*v1.OrdinaryRegisterReply,error) {
	//TODO to be completed
	return nil,nil
}