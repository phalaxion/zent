package store

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/phalaxion/zent/ledger"
)

type SQLiteStore struct {
	DB *sql.DB
}

func NewSQLiteStore(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Ensure the migrations table exists and apply provided migrations (if any).
	s := &SQLiteStore{DB: db}

	if err := s.ApplyMigrations(); err != nil {
		db.Close()
		return nil, err
	}

	return s, nil
}

func (s *SQLiteStore) RecordTransaction(t ledger.Transaction) error {
	stmt := `INSERT INTO transactions(id, amount, description, timestamp) VALUES(?, ?, ?, ?)`
	_, err := s.DB.Exec(stmt, t.ID, t.Amount, t.Description, t.Timestamp.UnixNano())
	return err
}

func (s *SQLiteStore) ListTransactions() ([]ledger.Transaction, error) {
	rows, err := s.DB.Query(`SELECT id, amount, description, timestamp FROM transactions ORDER BY timestamp`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []ledger.Transaction
	for rows.Next() {
		var id string
		var amount float64
		var description sql.NullString
		var ts int64
		if err := rows.Scan(&id, &amount, &description, &ts); err != nil {
			return nil, err
		}
		t := ledger.Transaction{
			ID:          id,
			Amount:      amount,
			Description: description.String,
			Timestamp:   time.Unix(0, ts),
		}
		out = append(out, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SQLiteStore) GetTransaction(id string) (*ledger.Transaction, error) {
	row := s.DB.QueryRow(`SELECT id, amount, description, timestamp FROM transactions WHERE id = ?`, id)
	var tid string
	var amount float64
	var description sql.NullString
	var ts int64
	if err := row.Scan(&tid, &amount, &description, &ts); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}
	t := &ledger.Transaction{
		ID:          tid,
		Amount:      amount,
		Description: description.String,
		Timestamp:   time.Unix(0, ts),
	}
	return t, nil
}

func (s *SQLiteStore) DeleteTransaction(id string) error {
	res, err := s.DB.Exec(`DELETE FROM transactions WHERE id = ?`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("transaction not found")
	}
	return nil
}

type Migration struct {
	ID      string
	Version int
	Up      string
}

func (s *SQLiteStore) ApplyMigrations() error {
	var currentVersion int
	err := s.DB.QueryRow("PRAGMA user_version").Scan(&currentVersion)
	if err != nil {
		return err
	}

	migrations := []Migration{}

	if currentVersion < 1 {
		migrations = append(migrations, Migration{
			ID:      "0001_create_transactions",
			Version: 1,
			Up: `CREATE TABLE IF NOT EXISTS transactions (
				id TEXT PRIMARY KEY,
				amount REAL NOT NULL,
				description TEXT,
				timestamp INTEGER NOT NULL
			);`,
		})
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	versionChanged := false
	for _, m := range migrations {
		if _, err := tx.Exec(m.Up); err != nil {
			tx.Rollback()
			return fmt.Errorf("migration %s failed: %w", m.ID, err)
		}

		if m.Version > 0 && m.Version > currentVersion {
			currentVersion = m.Version
			versionChanged = true
		}

	}

	if versionChanged {
		if _, err := tx.Exec(fmt.Sprintf("PRAGMA user_version = %d", currentVersion)); err != nil {
			tx.Rollback()
			return fmt.Errorf("setting user_version for %d failed: %w", currentVersion, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
