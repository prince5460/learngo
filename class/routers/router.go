package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"learngo/class/controllers"
)

func init() {
	beego.InsertFilter("/Article/*", beego.BeforeRouter, filterFunc)

	beego.Router("/", &controllers.MainController{})
	//用户
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	//文章
	beego.Router("/Article/index", &controllers.ArticleController{}, "get:ShowIndex")
	beego.Router("/Article/addArticle", &controllers.ArticleController{}, "get:ShowAdd;post:HandleAdd")
	beego.Router("/Article/content", &controllers.ArticleController{}, "get:ShowContent")
	beego.Router("/Article/update", &controllers.ArticleController{}, "get:ShowUpdate;post:HandleUpdate")
	beego.Router("/Article/delete", &controllers.ArticleController{}, "get:HandleDelete")
	//类型
	beego.Router("/Article/AddArticleType", &controllers.ArticleController{}, "get:ShowAddType;post:HandleAddType")
}

//过滤器函数,验证是否登录状态
var filterFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302, "/login")
	}
}
