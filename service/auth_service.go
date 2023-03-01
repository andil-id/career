package service

import (
	"career/model/web"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, req web.LoginRequest) (web.LoginResponse, error)
}
