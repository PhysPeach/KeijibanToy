package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int64 
	Name string `orm:"size(24)"`
	Text string `orm:"size(200)"`
	Created time.Time
}

func init() {
	fmt.Println("init models")
	orm.RegisterModel(new(Post))

	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlConn := "physpeach:@tcp(127.0.0.1:3306)/KeijibanToy?charset=utf8"
    if err := orm.RegisterDataBase("default", "mysql", mysqlConn); err != nil {
		panic(err)
	}
}

func GetAllPosts() []Post {
	o := orm.NewOrm()
	
	var posts []Post
    o.QueryTable(new(Post)).OrderBy("-created").All(&posts)

    return posts
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddPost(m *Post) () {
	o := orm.NewOrm()
	o.Insert(m)
	return
}