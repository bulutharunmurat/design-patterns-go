package main

type Invesment struct {
	initialAmount int
}

func (invesment *Invesment) deposit(amount int) {
	invesment.initialAmount = invesment.initialAmount + amount
}

func (invesment *Invesment) withdraw(amount int) {
	invesment.initialAmount = invesment.initialAmount - amount
}

func (invesment *Invesment) transfer(account IAccount, amount int) {
	invesment.withdraw(amount)
	account.deposit(amount)
}
func (invesment *Invesment) getAccountNumber() int {
	return 0
}
func (invesment *Invesment) getInitAmount() int {
	return invesment.initialAmount
}
