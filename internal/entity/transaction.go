package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.NewString(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	err := transaction.validate()
	if err != nil {
		return nil, err
	}
	transaction.Commit()
	return transaction, nil
}

func (t *Transaction) Commit() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}

func (t *Transaction) validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}
	return nil
}