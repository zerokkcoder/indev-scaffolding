package service

import (
	"github.com/zerokkcoder/indevsca/pkg/jwt"
	"github.com/zerokkcoder/indevsca/pkg/log"
)

// Service 服务
type Service struct {
	logger *log.Logger
	jwt    *jwt.JWT
}

// NewServer 创建服务
func NewServer(
	logger *log.Logger,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		jwt:    jwt,
	}
}
