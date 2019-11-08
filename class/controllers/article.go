package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"learngo/class/models"
	"math"
	"path"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

//处理下拉框请求
func (c *ArticleController) HandleSelect() {
	//1.接收数据
	typeName := c.GetString("select")
	//2.处理数据
	if typeName == "" {
		logs.Info("下拉框传递数据失败")
		return
	}
	//3.查询数据
	o := orm.NewOrm()
	var articles []models.Article
	//惰性查询,加上RelatedSel
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)

}

//显示文章首页
func (c *ArticleController) ShowIndex() {
	//判断是否登录
	userName := c.GetSession("userName")
	if userName == nil {
		c.Redirect("/login", 302)
		return
	}

	o := orm.NewOrm()
	var articles []models.Article
	//_, err := o.QueryTable("Article").All(&articles)
	//if err != nil {
	//	logs.Info("数据查询失败")
	//	return
	//}

	qs := o.QueryTable("Article")
	//_, err := qs.All(&articles) //select * from Article
	//if err != nil {
	//	logs.Info("数据查询失败")
	//	return
	//}

	//pageIndex := c.GetString("pageIndex")
	pageIndex, err := strconv.Atoi(c.GetString("pageIndex"))
	if err != nil {
		pageIndex = 1
	}

	count, err := qs.RelatedSel("ArticleType").Count() //返回数据条目数
	if err != nil {
		logs.Info("查询失败")
		return
	}
	pageSize := 2 //每页条目数

	start := pageSize * (pageIndex - 1) //起始位置
	_, err = qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articles)
	if err != nil {
		logs.Info("数据查询失败")
		return
	}
	pageCount := math.Ceil(float64(count) / float64(pageSize)) //总页数

	//首页末页数据处理
	FirstPage := false
	if pageIndex == 1 {
		FirstPage = true
	}
	LastPage := false
	if pageIndex == int(pageCount) {
		LastPage = true
	}

	//读取类型
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	//根据类型获取数据
	var articleWithType []models.Article
	typeName := c.GetString("select")
	if typeName == "" {
		logs.Info("下拉框传递数据失败")
		qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articleWithType)
	} else {
		qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articleWithType)
	}

	c.Data["userName"] = c.GetSession("userName")
	c.Data["types"] = types
	c.Data["FirstPage"] = FirstPage
	c.Data["LastPage"] = LastPage
	c.Data["count"] = count
	c.Data["pageCount"] = pageCount
	c.Data["pageIndex"] = pageIndex

	c.Data["articles"] = articles

	c.Layout = "layout.html"
	c.Data["msg"] = "文章列表页"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["contentHead"] = "head.html"
	c.TplName = "index.html"
}

//显示添加页面
func (c *ArticleController) ShowAdd() {
	//读取类型
	o := orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	c.Data["types"] = types
	c.Data["userName"] = c.GetSession("userName")
	c.Layout = "layout.html"
	c.Data["msg"] = "文章添加页"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["contentHead"] = "head.html"
	c.TplName = "add.html"
}

//判断slice中是否存在某个元素
func itemContains(s []string, item string) bool {
	for _, i := range s {
		if i == item {
			return true
		}
	}
	return false
}

//处理添加文章页面数据
func (c *ArticleController) HandleAdd() {
	//1.获取数据
	title := c.GetString("articleName")
	content := c.GetString("content")
	file, header, err := c.GetFile("uploadname")
	typename := c.GetString("select") //下拉框
	if typename == "" {
		logs.Info("下拉框数据错误")
		return
	}

	defer file.Close()

	//1.限定上传文件格式
	fileext := path.Ext(header.Filename)
	exts := []string{".jpg", ".png", ".jpeg", ".gif"}
	contains := itemContains(exts, fileext)
	if !contains {
		logs.Info("上传格式错误")
		return
	}
	//if fileext != ".jpg" && fileext != ".png" {
	//	logs.Info("上传格式错误")
	//	return
	//}
	//2.限制文件大小
	if header.Size > 5*1024*1024 {
		logs.Info("上传文件过大")
		return
	}
	//3.对文件重命名防止重复
	filename := time.Now().Format("20060102150405") + fileext

	if err != nil {
		logs.Info("上传文件失败")
		return
	} else {
		c.SaveToFile("uploadname", "../static/img/"+filename)
	}

	//2.判断数据是否合法
	if title == "" || content == "" {
		logs.Info("添加文章数据错误")
		return
	}

	//3.插入数据
	o := orm.NewOrm()
	article := models.Article{}
	article.Title = title
	article.Content = content
	article.Img = "./static/img/" + filename

	//获取type对象
	var articleType models.ArticleType
	articleType.TypeName = typename
	err = o.Read(&articleType, "TypeName")
	if err != nil {
		logs.Info("获取类型错误")
		return
	}
	article.ArticleType = &articleType

	_, err = o.Insert(&article)
	if err != nil {
		logs.Info("插入数据错误")
		return
	}

	//4.返回文章界面
	c.Redirect("/Article/index", 302)

}

//显示内容详情页
func (c *ArticleController) ShowContent() {
	//1.获取文章id
	id, err := c.GetInt("id")
	if err != nil {
		logs.Info("获取文章ID错误:", err)
		return
	}

	//2.查询数据库获取数据
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		logs.Info("查询错误:", err)
		return
	}
	//阅读次数
	article.Count += 1

	//多对多插入读者
	//1.获取操作对象
	//article := models.Article{Id: id}
	//2.获取多对多操作对象
	m2m := o.QueryM2M(&article, "Users")
	//3.获取插入对象
	userName := c.GetSession("userName")
	user := models.User{}
	user.UserName = userName.(string)
	o.Read(&user, "UserName")
	//4.多对多插入
	_, err = m2m.Add(&user)
	if err != nil {
		logs.Info("插入失败")
		return
	}
	o.Update(&article)

	//o.LoadRelated(&article,"Users")
	//o.QueryTable("Article").RelatedSel("User").Filter("Users__User__UserName",userName.(string)).Distinct().Filter("Id",id).One(&article)
	//开始读出历史浏览者信息
	users := []models.User{}
	o.QueryTable("User").Filter("Articles__Article__Id", article.Id).Distinct().All(&users)
	c.Data["users"] = users

	//3.传递数据给视图
	c.Data["article"] = article
	c.Data["userName"] = c.GetSession("userName")
	c.Layout = "layout.html"
	c.Data["msg"] = "文章详情页"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["contentHead"] = "head.html"
	c.TplName = "content.html"
}

//显示编辑页面
func (c *ArticleController) ShowUpdate() {
	//1.获取文章id
	id, err := c.GetInt("id")
	if err != nil {
		logs.Info("获取文章ID错误:", err)
		return
	}

	//2.查询数据库获取数据
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		logs.Info("查询错误:", err)
		return
	}

	//3.传递数据给视图
	c.Data["article"] = article
	c.Data["userName"] = c.GetSession("userName")
	c.Layout = "layout.html"
	c.Data["msg"] = "文章编辑页"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["contentHead"] = "head.html"
	c.TplName = "update.html"
}

//处理更新业务数据
func (c *ArticleController) HandleUpdate() {
	//1.获取数据
	id, _ := c.GetInt("id")
	title := c.GetString("articleName")
	content := c.GetString("content")
	file, header, err := c.GetFile("uploadname")
	var filename string
	if err != nil {
		logs.Info("上传文件失败")
		return
	} else {
		defer file.Close()
		//限定上传文件格式
		fileext := path.Ext(header.Filename)
		if fileext != ".jpg" && fileext != ".png" {
			logs.Info("上传格式错误")
			return
		}
		//限制文件大小
		if header.Size > 5*1024*1024 {
			logs.Info("上传文件过大")
			return
		}
		//对文件重命名防止重复
		filename = time.Now().Format("20060102150405") + fileext
		c.SaveToFile("uploadname", "../static/img/"+filename)
	}
	//2.处理数据
	if title == "" || content == "" {
		logs.Info("更新数据失败")
		return
	}

	//3.更新操作
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		logs.Info("查询数据错误")
		return
	}
	article.Title = title
	article.Content = content
	article.Img = "./static/img/" + filename
	_, err = o.Update(&article, "Title", "Content", "Img")
	if err != nil {
		logs.Info("更新数据错误")
		return
	}
	//4.返回列表页
	c.Redirect("/Article/index", 302)
}

//删除操作
func (c *ArticleController) HandleDelete() {
	//1.获取数据
	id, err := c.GetInt("id")
	if err != nil {
		logs.Info("获取id错误")
		return
	}

	//2.执行删除操作
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		logs.Info("查询数据错误")
		return
	}
	o.Delete(&article)

	//3.返回列表页
	c.Redirect("/Article/index", 302)

}

func (c *ArticleController) ShowAddType() {
	//1.读取类型表
	o := orm.NewOrm()
	var articleTypes []models.ArticleType
	//2.查询
	_, err := o.QueryTable("ArticleType").All(&articleTypes)
	if err != nil {
		logs.Info("查询类型错误")
		//return
	}
	c.Data["types"] = articleTypes
	c.Data["userName"] = c.GetSession("userName")
	c.Layout = "layout.html"
	c.Data["msg"] = "分类添加页"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["contentHead"] = "head.html"
	c.TplName = "addType.html"
}

//处理添加类型业务
func (c *ArticleController) HandleAddType() {
	//1.获取数据
	typename := c.GetString("typeName")
	//2.判断数据
	if typename == "" {
		logs.Info("添加类型不能为空")
		return
	}
	//3.执行插入数据
	o := orm.NewOrm()
	var articleType models.ArticleType
	articleType.TypeName = typename
	_, err := o.Insert(&articleType)
	if err != nil {
		logs.Info("添加失败")
		return
	}
	//4.展示视图
	c.Redirect("/Article/AddArticleType", 302)
}
