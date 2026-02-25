package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/phalaxion/zent/ledger"
)

type JSONStore struct {
	FilePath string
}

type ledgerFile struct {
	Transactions []ledger.Transaction `json:"transactions"`
}

func (s *JSONStore) RecordTransaction(t ledger.Transaction) error {
	lf, err := s.load()
	if err != nil {
		return err
	}

	lf.Transactions = append(lf.Transactions, t)
	return s.save(lf)
}

func (s *JSONStore) ListTransactions() ([]ledger.Transaction, error) {
	lf, err := s.load()
	if err != nil {
		return nil, err
	}
	return lf.Transactions, nil
}

func (s *JSONStore) GetTransaction(id string) (*ledger.Transaction, error) {
	lf, err := s.load()
	if err != nil {
		return nil, err
	}

	for _, transaction := range lf.Transactions {
		if transaction.ID == id {
			return &transaction, nil
		}
	}

	return nil, fmt.Errorf("transaction not found")
}

func (s *JSONStore) load() (*ledgerFile, error) {
	file, err := os.Open(s.FilePath)
	if errors.Is(err, os.ErrNotExist) {
		return &ledgerFile{Transactions: []ledger.Transaction{}}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lf ledgerFile
	if err := json.NewDecoder(file).Decode(&lf); err != nil {
		return nil, err
	}

	return &lf, nil
}

func (s *JSONStore) save(lf *ledgerFile) error {
	file, err := os.Create(s.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(lf)
}
