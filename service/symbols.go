package service

import (
	"context"
	"strings"

	"github.com/isther/easyTrader/model"
	"github.com/isther/easyTrader/pkg/account"
)

type SymbolsService struct{}

func NewSymbolsService() *SymbolsService {
	return &SymbolsService{}
}

func (s *SymbolsService) CheckIsOk(username, symbol string) (bool, error) {
	symbol = strings.ToUpper(symbol)

	uf, err := model.FindUserByUsername(username)
	if err != nil {
		return false, err
	}

	worker := &account.Worker{
		Context: context.Background(),
		Account: account.NewBinance(uf.Binance.ApiKey, uf.Binance.SecretKey),
	}

	return worker.CheckIsOK(symbol)
}

func (s *SymbolsService) SetSymbols(user *model.User, symbols ...string) error {
	return user.SetSymbols(symbols...)
}

func (s *SymbolsService) GetSymbols(username string) ([]string, error) {
	uf, err := model.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return uf.Symbols, nil
}

func (s *SymbolsService) SearchSymbols(username string, amplitue, amplituePercentage float64) ([]*account.SearchByAmplitueResult, error) {
	uf, err := model.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}

	worker := &account.Worker{
		Symbols: uf.Symbols,
		Context: context.Background(),
		Account: account.NewBinance(uf.Binance.ApiKey, uf.Binance.SecretKey),
	}

	return worker.SearchSymbolsByAmplitue(120, amplitue, amplituePercentage), nil
}
