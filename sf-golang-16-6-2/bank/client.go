package bank

import (
	"errors"
	"sync"
)

var (
	ErrBalance = errors.New("insufficient funds")
	ErrAmount  = errors.New("amount should not be negative")
)

// ClientInterface - интерфейс клиента банка
type ClientInterface interface {
	Deposit(amount int)
	Withdrawal(amount int) error
	Balance() int
}

// Client - потокобезопасная реализация ClientInterface
type Client struct {
	muBalance sync.RWMutex
	balance   int
}

func NewClient() *Client {
	return &Client{balance: 0}
}

// Withdrawal - списывает amount с баланса.
// Возвращает ошибку, если amount отрицательный
// или недостаточно средств на балансе
func (c *Client) Withdrawal(amount int) error {
	if amount < 0 {
		return ErrAmount
	}
	c.muBalance.Lock()
	defer c.muBalance.Unlock()
	if amount > c.balance {
		return ErrBalance
	}
	c.balance -= amount
	return nil
}

// Balance - возвращает баланс
func (c *Client) Balance() int {
	c.muBalance.RLock()
	b := c.balance
	c.muBalance.RUnlock()
	return b
}

// Deposit - пополняет баланс на значение amount.
// Вызывает панику, если amount отрицательный,
func (c *Client) Deposit(amount int) {
	if amount < 0 {
		panic(ErrAmount)
	}
	c.muBalance.Lock()
	c.balance += amount
	c.muBalance.Unlock()
}
