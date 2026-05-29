package main

import (
	"fmt"
	"strings"
	"sync"
)

type BankAccount struct {
	ID         string
	Owner      string
	Balance    float64
	MinBalance float64
	mu         sync.Mutex
}

const MaxTransactionAmount = 10000.0

type AccountError struct {
	Message string
	Code    string
}

func (e *AccountError) Error() string {
	return fmt.Sprintf("account error [%s]: %s", e.Code, e.Message)
}

type InsufficientFundsError struct {
	RequestedAmount float64
	CurrentBalance  float64
	MinBalance      float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: requested %.2f, current balance %.2f, minimum balance %.2f", e.RequestedAmount, e.CurrentBalance, e.MinBalance)
}

type NegativeAmountError struct {
	Amount float64
}

func (e *NegativeAmountError) Error() string {
	return fmt.Sprintf("negative amount not allowed: %.2f", e.Amount)
}

type ExceedsLimitError struct {
	Amount float64
	Limit  float64
}

func (e *ExceedsLimitError) Error() string {
	return fmt.Sprintf("amount %.2f exceeds transaction limit of %.2f", e.Amount, e.Limit)
}

func NewBankAccount(id, owner string, initialBalance, minBalance float64) (*BankAccount, error) {
	if strings.TrimSpace(id) == "" {
		return nil, &AccountError{
			Message: "account ID cannot be empty",
			Code:    "INVALID_ID",
		}
	}

	if strings.TrimSpace(owner) == "" {
		return nil, &AccountError{
			Message: "account owner cannot be empty",
			Code:    "INVALID_OWNER",
		}
	}

	if initialBalance < 0 {
		return nil, &NegativeAmountError{Amount: initialBalance}
	}

	if minBalance < 0 {
		return nil, &NegativeAmountError{Amount: minBalance}
	}

	if initialBalance < minBalance {
		return nil, &InsufficientFundsError{
			RequestedAmount: 0,
			CurrentBalance:  initialBalance,
			MinBalance:      minBalance,
		}
	}

	return &BankAccount{
		ID:         strings.TrimSpace(id),
		Owner:      strings.TrimSpace(owner),
		Balance:    initialBalance,
		MinBalance: minBalance,
	}, nil
}

func (a *BankAccount) Deposit(amount float64) error {

	if amount < 0 {
		return &NegativeAmountError{Amount: amount}
	}

	if amount > MaxTransactionAmount {
		return &ExceedsLimitError{
			Amount: amount,
			Limit:  MaxTransactionAmount,
		}
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += amount
	return nil
}

func (a *BankAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return &NegativeAmountError{Amount: amount}
	}

	if amount > MaxTransactionAmount {
		return &ExceedsLimitError{
			Amount: amount,
			Limit:  MaxTransactionAmount,
		}
	}

	if amount >= 0 && a.Balance-amount < a.MinBalance {
		return &InsufficientFundsError{
			RequestedAmount: amount,
			CurrentBalance:  a.Balance,
			MinBalance:      a.MinBalance,
		}
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance -= amount

	return nil
}

func (a *BankAccount) Transfer(amount float64, target *BankAccount) error {

	if target == nil {
		return &AccountError{
			Message: "target account cannot be nil",
			Code:    "INVALID_TARGET",
		}
	}

	if amount < 0 {
		return &NegativeAmountError{Amount: amount}
	}

	if amount > MaxTransactionAmount {
		return &ExceedsLimitError{
			Amount: amount,
			Limit:  MaxTransactionAmount,
		}
	}

	if a == target {
		return &AccountError{
			Message: "cannot transfer to the same account",
			Code:    "SELF_TRANSFER",
		}
	}

	if amount >= 0 && a.Balance-amount < a.MinBalance {
		return &InsufficientFundsError{
			RequestedAmount: amount,
			CurrentBalance:  a.Balance,
			MinBalance:      a.MinBalance,
		}
	}

	var first, second *BankAccount
	if strings.Compare(a.ID, target.ID) < 0 {
		first, second = a, target
	} else {
		first, second = target, a
	}

	first.mu.Lock()
	defer first.mu.Unlock()

	second.mu.Lock()
	defer second.mu.Unlock()

	a.Balance -= amount
	target.Balance += amount

	return nil
}
