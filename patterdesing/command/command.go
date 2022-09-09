package main

import (
	"fmt"
)

var overdraftLimit = -500

type BankAcount struct {
	balance int
}

func (b *BankAcount) Deposit(amount int) {
	b.balance += amount
}

func (b *BankAcount) With(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		return true
	}

	return false
}

type Command interface {
	call()
	undo()
	Succeeded() bool
	SetSucceeded(value bool)
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	acoount   *BankAcount
	action    Action
	amount    int
	succeeded bool
}

func (b *BankAccountCommand) Succeeded() bool {
	return b.succeeded
}

func (b *BankAccountCommand) SetSucceeded(value bool) {
	b.succeeded = value
}

func (b *BankAccountCommand) call() {
	switch b.action {
	case Deposit:
		b.acoount.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.acoount.With(b.amount)
	}

	fmt.Printf("total amount %v \n", b.acoount.balance)
}

func (b *BankAccountCommand) undo() {
	if !b.succeeded {
		return
	}
	switch b.action {
	case Deposit:
		b.acoount.With(b.amount)
	case Withdraw:
		b.acoount.Deposit(b.amount)
	}
}

type CompositeBank struct {
	commands []Command
}

func (c CompositeBank) call() {
	for _, cmd := range c.commands {
		cmd.call()
	}
}

func (c CompositeBank) undo() {
	for idx := range c.commands {
		c.commands[len(c.commands)-idx-1].undo()
	}
}

func (c CompositeBank) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}

	return true
}

func (c CompositeBank) SetSucceeded(value bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(value)
	}
}

func NewBankAccountCommand(account *BankAcount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{
		acoount: account,
		action:  action,
		amount:  amount,
	}
}

type MoneyTransferCommand struct {
	CompositeBank
	from, to *BankAcount
	amount   int
}

func (m *MoneyTransferCommand) call() {
	ok := true

	for _, cmd := range m.commands {
		if ok {
			cmd.call()
			ok = cmd.Succeeded()
			continue
		}
		cmd.SetSucceeded(false)
	}
}

func NewMoneyTrasferCommand(from *BankAcount, to *BankAcount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands, NewBankAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands, NewBankAccountCommand(to, Deposit, amount))

	return c
}

func main() {
	ba := &BankAcount{}
	cmd := NewBankAccountCommand(ba, Deposit, 100)
	cmd.call()

	cmd2 := NewBankAccountCommand(ba, Withdraw, 50)
	cmd2.call()
	cmd2.undo()
	fmt.Println(ba.balance, "new total")
	fmt.Println("new balance to")
	from := &BankAcount{100}
	to := &BankAcount{0}

	mtc := NewMoneyTrasferCommand(from, to, 25)
	mtc.call()
}
