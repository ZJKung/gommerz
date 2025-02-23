package account

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]*Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	repo Repository
}

func newService(r Repository) *accountService {
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	account := &Account{
		ID:   uuid.New().String(),
		Name: name,
	}

	err := s.repo.PubAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return s.repo.GetAccountByID(ctx, id)
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]*Account, error) {
	if take > 100 || (take == 0 && skip == 0) {
		take = 100
	}
	return s.repo.ListAccounts(ctx, skip, take)
}
