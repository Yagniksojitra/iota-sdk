package expense

import (
	"github.com/iota-uz/iota-sdk/modules/finance/domain/aggregates/expense_category"
	moneyaccount "github.com/iota-uz/iota-sdk/modules/finance/domain/aggregates/money_account"
	"time"
)

type Expense struct {
	ID               uint
	Amount           float64
	Account          moneyaccount.Account
	Category         category.ExpenseCategory
	Comment          string
	TransactionID    uint
	AccountingPeriod time.Time
	Date             time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
