package service

import (
	"context"
	"errors"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/model"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Password != req.ConfirmPassword {
		err = errors.New("password must be the same as ConfirmPassword")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
	}
	if err = model.Create(mysql.DB, s.ctx, newUser); err != nil {
		return
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
