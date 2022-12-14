package service

import (
	"errors"

	"github.com/isther/easyTrader/model"
	"github.com/isther/easyTrader/pkg/jwt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(user *model.User) error {
	if model.UsernameExist(user.Username) {
		return errors.New("username exist")
	}
	return model.NewUser(user)
}

func (u *UserService) Login(user *model.User) (string, error) {
	tmpUser, err := model.FindUserByUsername(user.Username)
	if err != nil {
		return "", err
	}

	if user.Password != tmpUser.Password {
		return "", errors.New("error password")
	}

	token, err := jwt.BuildMapClaimsJwt(user.Username, user.Password)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserService) SetPassword(user *model.User, password string) error {
	return user.SetPassword(password)
}

func (u *UserService) SetBinanceApiKey(user *model.User, apiKey string) error {
	return user.SetBinanceApiKey(apiKey)
}

func (u *UserService) SetBinanceSecretKey(user *model.User, secretKey string) error {
	return user.SetBinanceSecretKey(secretKey)
}

func (u *UserService) SetDingDingTalkAccessToken(user *model.User, accessToken string) error {
	return user.SetDingDingTalkApiKey(accessToken)
}

func (u *UserService) SetDingDingTalkSecret(user *model.User, secret string) error {
	return user.SetDingDingTalkSecretKey(secret)
}

func (u *UserService) SetDingDingTalkEnable(user *model.User, enable bool) error {
	return user.EnableDingDingTalk(enable)
}

func (u *UserService) GetBinanceApiKey(username string) (string, error) {
	uf, err := model.FindUserByUsername(username)
	if err != nil {
		return "", err
	}

	return uf.Binance.ApiKey, nil
}

func (u *UserService) GetBinanceSecretKey(username string) (string, error) {
	// uf, err := model.FindUserByUsername(username)
	// if err != nil {
	// return "", err
	// }

	// return uf.Binance.SecretKey, nil
	return "", nil
}

func (u *UserService) GetDingDingTalkAccessToken(username string) (string, error) {
	uf, err := model.FindUserByUsername(username)
	if err != nil {
		return "", err
	}

	return uf.DingDingTalk.AccessToken, nil
}

func (u *UserService) GetDingDingTalkSecret(username string) (string, error) {
	// uf, err := model.FindUserByUsername(username)
	// if err != nil {
	// return "", err
	// }

	// return uf.DingDingTalk.Secret, nil
	return "", nil
}

func (u *UserService) GetDingDingTalkEnable(username string) (bool, error) {
	uf, err := model.FindUserByUsername(username)
	if err != nil {
		return false, err
	}

	return uf.DingDingTalk.Enable, nil
}
