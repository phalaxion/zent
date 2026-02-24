package format

import "fmt"

func Currency(amount float64) string {
	if amount < 0 {
		return fmt.Sprintf("-$%.2f", -amount)
	}
	return fmt.Sprintf("$%.2f", amount)
}
