package main

import (
	"fmt"
	"jikeblog/models/class"
	_ "jikeblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//func init() {
//	fmt.Println(9876)
//	gob.Register(class.User{})
//}

func init() {
	fmt.Println(88888)
	switch beego.AppConfig.String("DB::db") {
	case "mysql":
		orm.RegisterDataBase("default", "mysql", "root:123456@/jblog?charset=utf8", 30)
		//orm.RegisterDriver("mysql", orm.DR_MySQL)
		//		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&loc=%s",
		//			beego.AppConfig.String("DB::user"),
		//			beego.AppConfig.String("DB::pass"),
		//			beego.AppConfig.String("DB::name"),
		//			`Asia%2FShanghai`,
		//		))
		//orm.RunSyncdb("default", false, true)
		//	case "sqlite":
		//		orm.RegisterDriver("sqlite", orm.DR_Sqlite)
		//		orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("DB::file"))
	}

	orm.RegisterModel(new(class.User), new(class.Article))

	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
