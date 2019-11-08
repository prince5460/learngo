package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//显示注册页面
func (c *MainController) Get() {
	/** 插入数据
	//1.新建orm对象
	o := orm.NewOrm()
	//2.一个要插入对象的结构体数据
	user := models.User{}
	//3.对结构体赋值
	user.Name = "Hello"
	user.Pwd = "123"
	//4.插入数据
	_, err := o.Insert(&user)
	if err != nil {
		logs.Info("插入数据失败,err = ", err)
	}
	*/

	/** 查询数据
	//1.有orm对象
	o := orm.NewOrm()
	//2.查询的对象
	user := models.User{}
	//3.指定查询对象字段值
	//user.Id = 5
	user.Name = "hello"
	//4.查询
	//err := o.Read(&user)
	err := o.Read(&user, "Name")
	if err != nil {
		logs.Info("查询失败,err = ", err)
	}
	logs.Info("查询成功", user)
	*/

	/** 更新数据
	//1.有orm对象
		o := orm.NewOrm()
		//2.需要更新的结构体对象
		user := models.User{}
		//3.查找需要更新的数据
		user.Id = 1
		err := o.Read(&user)
		//4.给数据重新赋值
		if err != nil {
			logs.Info("没有找到数据,err = ", err)
		}
		user.Name = "tom"
		user.Pwd = "123456"
		//5.更新
		_,err = o.Update(&user)
		if err != nil {
			logs.Info("更新失败,err = ", err)
			return
		}
		logs.Info("更新成功", user)*/

	/* 删除数据
	   //1.有orm对象
	   	o := orm.NewOrm()
	   	//2.删除的对象
	   	user := models.User{}
	   	//3.指定删除的数据
	   	user.Id = 1
	   	//4.删除
	   	_, err := o.Delete(&user)
	   	if err != nil {
	   		logs.Info("删除失败,err = ", err)
	   		return
	   	}
	   	logs.Info("删除成功", user)*/

	//c.Data["data"] = "今天天气不错"
	c.TplName = "index.tpl"
}
