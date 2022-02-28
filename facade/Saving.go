package main

type Saving struct {
	initialAmount int
}

func (saving *Saving) deposit(amount int) {
	saving.initialAmount = saving.initialAmount + amount
}

func (saving *Saving) withdraw(amount int) {
	saving.initialAmount = saving.initialAmount - amount
}

func (saving *Saving) transfer(account IAccount, amount int) {
	saving.withdraw(amount)
	account.deposit(amount)
}

func (saving *Saving) getAccountNumber() int {
	return 1
}
func (saving *Saving) getInitAmount() int {
	return saving.initialAmount
}
