## Helpful Commands

Build the app

```
go build -o zent main.go
```

## Goals!

💡 CLI / UX Improvements

- Add a delete command – remove transactions by ID. ✔️
- Add an edit command – modify amount, description, or timestamp for a given ID.
- Filter transactions – by date range, description keyword, or amount range.
- Search transactions – keyword search in descriptions.
- Pretty-print ledger – align columns, add totals, optionally color negative
  amounts red.
- Pagination – for list if there are many transactions.
- Export / Import CSV – allow exporting or importing transactions.
- Undo last transaction – a simple undo command.

🏦 Accounting / Business Logic

- Add categories – e.g., Food, Bills, Entertainment.
- Category summary – show balance or spending per category.
- Recurring transactions – auto-add transactions like subscriptions.
- Balance over time – show running balance per day/week/month.
- Budget tracking – allow user to set budgets and track against them.

🗄️ Storage / Architecture

- Switch to SQLite – more scalable, allows queries, avoids reading/writing full
  JSON.
- Move ledger file to user config directory – e.g., ~/.ledger/ledger.json.
- Backup / version history – keep previous versions of ledger in case of
  mistakes.
- Encryption / password protection – store sensitive financial data securely.

🧪 Testing / Quality

- Unit tests for service layer – test Add, List, Balance, Delete, etc.
- Mock stores for testing – simulate different storage backends without touching
  files.

Bonus “Fun / Advanced”

- Generate charts (ASCII or using termdash) for spending trends.
- Add multi-currency support.
- CLI auto-completion for commands and IDs.
