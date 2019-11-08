package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"learngo/class/models"
	"time"
)

type UserController struct {
	beego.Controller
}

//显示注册页面
func (c *UserController) ShowRegister() {
	c.TplName = "register.html"
}

//处理注册逻辑
func (c *UserController) HandleRegister() {
	//1.获取数据
	username := c.GetString("userName")
	password := c.GetString("password")
	//2.对数据进行校验
	if username == "" || password == "" {
		//c.Redirect("/register", 302)
		c.TplName = "register.html"
		c.Data["errmsg"] = "数据不完整，请检查！"
		return
	}
	//3.插入数据
	o := orm.NewOrm()
	user := models.User{
		UserName: username,
		PassWord: password,
	}

	_, err := o.Insert(&user)
	if err != nil {
		//c.Redirect("/register", 302)
		c.TplName = "register.html"
		c.Data["errmsg"] = "注册失败，请尝试换个用户名注册！"
		return
	}
	//4.返回登录界面
	c.Redirect("/login", 302)

}

//显示登录页面
func (c *UserController) ShowLogin() {
	username := c.Ctx.GetCookie("userName")
	if username != "" {
		c.Data["userName"] = username
		c.Data["check"] = "checked"
	}
	c.TplName = "login.html"
}

//处理登录逻辑
func (c *UserController) HandleLogin() {
	//1.获取数据
	username := c.GetString("userName")
	password := c.GetString("password")
	//2.对数据进行校验
	if username == "" || password == "" {
		//logs.Info("数据不能为空")
		//c.Redirect("/login", 302)
		c.Data["errmsg"] = "登录信息不完整！"
		c.TplName = "login.html"
		return
	}
	//3.查询数据
	o := orm.NewOrm()
	user := models.User{}
	user.UserName = username
	//user.PassWord = password

	err := o.Read(&user, "UserName")
	if err != nil {
		//logs.Info("登录失败:", err)
		c.Data["errmsg"] = "登录失败"
		c.TplName = "login.html"
		return
	}
	//判断秘密是否一致
	if user.PassWord != password {
		c.Data["errmsg"] = "密码错误"
		//logs.Info("密码错误:", err)
		c.TplName = "login.html"
		return
	}

	//记住用户名
	check := c.GetString("remember")
	if check == "on" {
		c.Ctx.SetCookie("userName", username, time.Second*3600)
	} else {
		c.Ctx.SetCookie("userName", "", -1)
	}

	//设置session
	c.SetSession("userName", username)

	//4.登录成功
	c.Redirect("/Article/index", 302)
}

//退出登录
func (c *UserController) Logout() {
	//1.删除登录状态
	c.DelSession("userName")
	//2.调整登录页面
	c.Redirect("/login", 302)

}
