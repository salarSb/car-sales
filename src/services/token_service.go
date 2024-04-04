package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/constants"
	"github.com/salarSb/car-sales/pkg/logging"
	"github.com/salarSb/car-sales/pkg/service_errors"
	"time"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)
	return &TokenService{logger: logger, cfg: cfg}
}

func (s *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(s.cfg.Jwt.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(s.cfg.Jwt.RefreshTokenExpireDuration * time.Minute).Unix()
	atc := jwt.MapClaims{}
	atc[constants.UserIdKey] = token.UserId
	atc[constants.FirstNameKey] = token.FirstName
	atc[constants.LastNameKey] = token.LastName
	atc[constants.UsernameKey] = token.Username
	atc[constants.MobileNumberKey] = token.MobileNumber
	atc[constants.EmailKey] = token.Email
	atc[constants.RolesKey] = token.Roles
	atc[constants.ExpireTimeKey] = td.AccessTokenExpireTime
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	var err error
	td.AccessToken, err = at.SignedString([]byte(s.cfg.Jwt.Secret))
	if err != nil {
		return nil, err
	}
	rtc := jwt.MapClaims{}
	rtc[constants.UserIdKey] = token.UserId
	rtc[constants.ExpireTimeKey] = td.RefreshTokenExpireTime
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)
	td.RefreshToken, err = rt.SignedString([]byte(s.cfg.Jwt.RefreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnexpectedError}
		}
		return []byte(s.cfg.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}
	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}
