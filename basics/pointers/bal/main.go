package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?
func updateBalance(c *customer, t transaction) error{
	switch t.transactionType {
		case transactionWithdrawal:
			if c.balance < t.amount {
				err := errors.New("insufficient funds")
				return err
			}else{
				c.balance -= t.amount
				return nil
			}
		case transactionDeposit:
			c.balance += t.amount
			return nil
		default:
			err := errors.New("unknown transaction type")
			return err
	}

}