package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"learngo/class/conf"
	"time"
)

type User struct {
	Id       int    `orm:"pk;auto"`
	UserName string `orm:"unique"`
	PassWord string
	Articles []*Article `orm:"rel(m2m)"` //多对多,会创建关系表
}

type Article struct {
	Id          int          `orm:"pk;auto"`
	Title       string       `orm:"size(20)"`
	Content     string       `orm:"size(500)"`
	Img         string       `orm:"size(50);null"`
	Time        time.Time    `orm:"auto_now_add;type(datetime)"`
	Count       int          `orm:"default(0)"`
	ArticleType *ArticleType `orm:"rel(fk)"`       //外键
	Users       []*User      `orm:"reverse(many)"` //多对多
}

type ArticleType struct {
	Id       int
	TypeName string     `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"` //一对多
}

func init() {
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", conf.DataSource)
	//映射model数据
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	//生成表
	orm.RunSyncdb("default", false, true)
}
