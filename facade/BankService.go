package main

import (
	"fmt"
)

type BankService struct {
	bankAccounts map[int]IAccount
}

func (b *BankService) createNewAccount(accountType string, amount int) IAccount {

	var newAccount IAccount

	switch accountType {
	case "saving":
		newAccount = &Saving{initialAmount: amount}
	case "invesment":
		newAccount = &Invesment{initialAmount: amount}
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

func (b *BankService) transferMoney(to int, from int, amount int) {
	accountMap := b.bankAccounts
	toAccount := accountMap[to]
	fromAccount := accountMap[from]

	fromAccount.getInitAmount()
	fromAccount.transfer(toAccount, amount)

}

func main() {

	service := BankService{bankAccounts: map[int]IAccount{}}

	saving := service.createNewAccount("saving", 100)
	fmt.Println("saving initial amount is :", saving.getInitAmount())

	invesment := service.createNewAccount("invesment", 500)
	fmt.Println("invesment initial amount is :", invesment.getInitAmount())

	fmt.Println("25 is transferred from saving to invesment")
	service.transferMoney(0, 1, 25)

	fmt.Println("after transfer saving initial amount is :", saving.getInitAmount())
	fmt.Println("after transfer invesment initial amount is :", invesment.getInitAmount())

}
