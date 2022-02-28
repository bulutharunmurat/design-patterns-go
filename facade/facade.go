package main

import (
	"fmt"
	"log"
)

type IAccount interface {
	deposit(amount int)
	withdraw(amount int)
	transfer(account IAccount, amount int)
	getAccountNumber() int
	getInitAmount() int
}

type Invesment struct {
	initialAmount int
}

type Saving struct {
	initialAmount int
}

func (invesment Invesment) deposit(amount int) {
}

func (invesment Invesment) withdraw(amount int) {
}

func (invesment Invesment) transfer(account IAccount, amount int) {
}
func (invesment Invesment) getAccountNumber() int {
	return 0
}
func (invesment Invesment) getInitAmount() int {
	return invesment.initialAmount
}

func (saving Saving) deposit(amount int) {
}

func (saving Saving) withdraw(amount int) {
}

func (saving Saving) transfer(account IAccount, amount int) {
}
func (saving Saving) getAccountNumber() int {
	return 1
}
func (saving Saving) getInitAmount() int {
	return saving.initialAmount
}

type BankService struct {
	bankAccounts map[int]IAccount
}

func (b BankService) createNewAccount(accountType string, amount int) IAccount {

	var newAccount IAccount

	switch accountType {
	case "saving":
		newAccount = Saving{initialAmount: amount}
	case "invesment":
		newAccount = Invesment{initialAmount: amount}
	default:
		fmt.Printf("invalid account type")
		return nil
	}
	if newAccount != nil {
		accountMap := b.bankAccounts
		accountMap[newAccount.getAccountNumber()] = newAccount
		return newAccount
	}
	return nil
}

func (b BankService) transferMoney(to int, from int, amount int) {
	accountMap := b.bankAccounts
	toAccount := accountMap[to]
	fromAccount := accountMap[from]
	log.Println("money is transfred from to", toAccount, fromAccount)
	fromAccount.getInitAmount()
	fromAccount.transfer(toAccount, 50)

}

func main() {

	service := BankService{bankAccounts: map[int]IAccount{}}

	saving := service.createNewAccount("saving", 100)
	accountNumberofSaving := saving.getAccountNumber()
	fmt.Println(accountNumberofSaving)
	fmt.Println(saving.getInitAmount())

	invesment := service.createNewAccount("invesment", 500)
	accountNumberofInvesment := invesment.getAccountNumber()
	fmt.Println(accountNumberofInvesment)
	fmt.Println(invesment.getInitAmount())

	service.transferMoney(0, 1, 50)

	fmt.Println(saving.getInitAmount())
	fmt.Println(invesment.getInitAmount())

}
