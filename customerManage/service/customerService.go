package service

import (
	"learngo/customerManage/model"
)

//完成对Customer的操作

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.InitCustomer(1, "zhangsan", "male", 20, "188", "123@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService

}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := range this.customers {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

//修改指定客户信息
func (this *CustomerService) Update(customer model.Customer) bool {
	index := this.FindById(customer.Id)
	if index == -1 {
		return false
	}
	this.customers = append(append(this.customers[:index], customer), this.customers[index+1:]...)
	return true
}
