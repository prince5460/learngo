package main

import (
	"fmt"
)

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Save(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误！")
		return
	}

	if money <= 0 {
		fmt.Println("输入金额有误！")
		return
	}

	account.Balance += money
	fmt.Println("存款成功！")
}

func (account *Account) Draw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误！")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("输入金额有误！")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功！")
}

func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误！")
		return
	}

	fmt.Printf("Account= %s,Balance = %f\n", account.AccountNo, account.Balance)

}

func main() {
	account := Account{
		AccountNo: "NO:123456",
		Pwd:       "000000",
		Balance:   100.0,
	}
	account.Query("000000")
	account.Save(200.0, "000000")
	account.Query("000000")
	account.Draw(50.0, "000000")
	account.Query("000000")

}
