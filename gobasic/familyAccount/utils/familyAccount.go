package utils

import (
	"fmt"
)

type familyAccount struct {
	//声明必要字段
	balance float64
	money   float64
	key     string
	note    string
	details string
	loop    bool
	flag    bool
}

func NewFamilyAccount() *familyAccount {
	//工厂模式的构造方法
	return &familyAccount{
		balance: 100.0,
		money:   0.0,
		key:     "",
		note:    "",
		details: "收支\t账户金额\t收支金额\t说明",
		loop:    true,
		flag:    false,
	}
}

func (account *familyAccount) showDetails() {
	fmt.Println("------当前收支明细记录------")
	if account.flag {
		fmt.Println(account.details)
	} else {
		fmt.Println("当前没有明细！")
	}
}

func (account *familyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&account.money)
	account.balance += account.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&account.note)
	account.details += fmt.Sprintf("\n收入\t%v\t%v\t%s", account.balance, account.money, account.note)
	account.flag = true
}

func (account *familyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&account.money)

	if account.money > account.balance {
		fmt.Println("余额不足！")
	}

	account.balance -= account.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&account.note)
	account.details += fmt.Sprintf("\n支出\t%v\t%v\t%s", account.balance, account.money, account.note)
	account.flag = true
}

func (account *familyAccount) exit() {
	fmt.Println("确定要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("输入有误！")
	}
	if choice == "y" {
		account.loop = false
	}
}

func (account *familyAccount) MainMenu() {
	for {
		fmt.Println("\n------家庭收支记账软件------")
		fmt.Println(" 	1.收支明细")
		fmt.Println(" 	2.登记收入")
		fmt.Println(" 	3.登记支出")
		fmt.Println(" 	4.退出软件")
		fmt.Println("请选择（1-4）:")
		fmt.Scanln(&account.key)
		switch account.key {
		case "1":
			account.showDetails()
		case "2":
			account.income()
		case "3":
			account.pay()
		case "4":
			account.exit()

		default:
			fmt.Println("输入有误！")
		}

		if !account.loop {
			break
		}

	}
}
