package ledger

import (
	"time"

	"github.com/google/uuid"
)

type Store interface {
	RecordTransaction(t Transaction) error
	ListTransactions() ([]Transaction, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) Add(amount float64, description string) error {
	t := Transaction{
		ID:          uuid.NewString(),
		Amount:      amount,
		Description: description,
		Timestamp:   time.Now(),
	}

	return s.store.RecordTransaction(t)
}

func (s *Service) List() ([]Transaction, error) {
	return s.store.ListTransactions()
}

func (s *Service) Balance() (float64, error) {
	transactions, err := s.store.ListTransactions()
	if err != nil {
		return 0, err
	}

	var balance float64
	for _, t := range transactions {
		balance += t.Amount
	}

	return balance, nil
}
