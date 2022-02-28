package main

type IAccount interface {
	deposit(amount int)
	withdraw(amount int)
	transfer(account IAccount, amount int)
	getAccountNumber() int
	getInitAmount() int
}
