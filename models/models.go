package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"

	//_ "github.com/go-sql-driver/mysql"曾经使用过mysql

	"os"
	"path"
	"strconv"
	"time"
)

const (
	//设置数据库路径
	_DB_NAME = "data/beeblog.db"
	//设置数据库名称
	_SQLITE3_DRIVER = "sqlite3"

	//_MYSQL_DRIVER = "mysql"
)

//分类
type Category struct {
	Id            int64
	Title         string
	Cteated       time.Time `orm:"index"`
	Views         int64     `orm:"index"`
	TopicTime     time.Time `orm:"index"`
	TopicCount    int64
	TopicLastUser int64
}

//文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Cteated         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	//检查数据库模型
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	//注册模型
	orm.RegisterModel(new(Category), new(Topic))
	//注册驱动（“SQLite3”属于默认注册，此处代码可省略）
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	//注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name}

	qs := o.QueryTable("Category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("Category")
	_, err := qs.All(&cates)

	return cates, err
}
