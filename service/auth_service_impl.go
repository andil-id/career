package service

import (
	"career/exception"
	"career/helper"
	"career/model/web"
	"career/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	e "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	DB              *sql.DB
	Validate        *validator.Validate
	AdminRepository repository.AdminRepository
}

func NewAuthService(db *sql.DB, validate *validator.Validate, adminRepository repository.AdminRepository) AuthService {
	return &AuthServiceImpl{
		DB:              db,
		Validate:        validate,
		AdminRepository: adminRepository,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, req web.LoginRequest) (web.LoginResponse, error) {
	res := web.LoginResponse{}

	err := s.Validate.Struct(req)
	if err != nil {
		return res, err
	}

	admin, err := s.AdminRepository.GetAdminByUsername(ctx, s.DB, req.Username)
	if err != nil {
		return res, e.Wrap(exception.ErrBadRequest, "Wrong username")
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return res, e.Wrap(exception.ErrBadRequest, "Wrong password")
	}
	signedToken, err := helper.GenereateJwtToken(admin.Id, admin.Username)
	if err != nil {
		return res, err
	}

	res = web.LoginResponse{
		Id:        admin.Id,
		Username:  admin.Username,
		Token:     signedToken,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
	return res, nil
}
