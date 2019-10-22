package view

import (
	"fmt"
	"learngo/customerManage/model"
	"learngo/customerManage/service"
)

type customerView struct {
	key  string //接收用户输入
	loop bool   //表示是否循环显示主菜单

	customerService *service.CustomerService
}

func NewcustomerView() *customerView {
	return &customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}
}

//显示用户信息
func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("------用户列表------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := range customers {
		fmt.Println(customers[i].GetInfo())
	}

	fmt.Println("\n------用户列表完成------")

}

func (this *customerView) add() {
	fmt.Println("------添加用户------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer(name, gender, age, phone, email)

	if this.customerService.Add(customer) {
		fmt.Println("------添加完成------")
	} else {
		fmt.Println("------添加失败------")
	}
}

func (this *customerView) delete() {
	fmt.Println("------删除用户------")
	fmt.Println("请选择要删除的用户编号（-1退出）:")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除（Y/N）:")

	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("------删除成功------")
		} else {
			fmt.Println("------删除失败------")
		}
	}
}

//得到用户的输入，修改对应Id的客户
func (this *customerView) update() {
	fmt.Println("----------修改客户----------")
	fmt.Print("请选择修改客户的编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("电邮：")
	email := ""
	fmt.Scanln(&email)

	customer := model.InitCustomer(id, name, gender, age, phone, email)

	if this.customerService.Update(customer) {
		fmt.Println("----------修改成功----------")
	} else {
		fmt.Println("----------修改失败----------")
	}
}

func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("输入有误，确认是否退出(Y/N)")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}
}

//显示主菜单
func (this *customerView) MainMenu() {
	for {
		fmt.Println("------用户信息管理------")
		fmt.Println("1.添加用户")
		fmt.Println("2.修改用户")
		fmt.Println("3.删除用户")
		fmt.Println("4.用户列表")
		fmt.Println("5.退出")
		fmt.Println("请选择(1-5):")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()

		default:
			fmt.Println("输入有误，请重新输入!")
		}

		if !this.loop {
			break
		}
	}

	fmt.Println("退出成功！")
}
