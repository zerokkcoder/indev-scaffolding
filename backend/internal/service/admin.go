package service

import (
	"context"
	"time"

	"github.com/zerokkcoder/indevsca/internal/model"
	"github.com/zerokkcoder/indevsca/internal/repository"
	"github.com/zerokkcoder/indevsca/internal/request/admin"
	"github.com/zerokkcoder/indevsca/pkg/reponse"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	Login(ctx context.Context, req *admin.LoginRequest) (*model.Admin, string, error)
}

func NewAdminService(
	service *Service,
	adminRepo repository.AdminRepository,
) AdminService {
	return &adminService{
		service:   service,
		adminRepo: adminRepo,
	}
}

type adminService struct {
	service   *Service
	adminRepo repository.AdminRepository
}

func (s *adminService) Login(ctx context.Context, req *admin.LoginRequest) (*model.Admin, string, error) {
	admin, err := s.adminRepo.GetByUserName(ctx, req.Username)
	if err != nil || admin == nil {
		return nil, "", reponse.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return nil, "", err
	}

	token, err := s.service.jwt.GenToken(admin.ID, time.Now().Add(time.Hour*24*90))

	return admin, token, nil
}
