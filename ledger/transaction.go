package ledger

import "time"

type Transaction struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}
