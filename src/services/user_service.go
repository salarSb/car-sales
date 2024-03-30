package services

import (
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/common"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/db"
	"github.com/salarSb/car-sales/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{cfg: cfg, database: database, logger: logger, otpService: NewOtpService(cfg)}
}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
