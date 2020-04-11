// Beego Golang API using the Standard Package Layout.
package main

import (
	groot "beego-standard-layout"
	"beego-standard-layout/controllers"
	_ "beego-standard-layout/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

// Setup database for the first time when the app starts.
func setupDB() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		panic(err)
	}

	dbHost := beego.AppConfig.String("DataSourceHost")
	dbPort := beego.AppConfig.String("DataSourcePort")
	dbUser := beego.AppConfig.String("DataSourceUser")
	dbPass := beego.AppConfig.String("DataSourcePassword")
	dbName := beego.AppConfig.String("DataSourceDBName")
	dbSchema := beego.AppConfig.String("DataSourceSchema")

	dsPath := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable search_path=%s",
		dbUser, dbPass, dbName, dbHost, dbPort, dbSchema)

	err = orm.RegisterDataBase("default", "postgres", dsPath)
	if err != nil {
		panic(err)
	}

	orm.RegisterModel(new(groot.Struct))
	orm.Debug = true
}

func main() {
	r := beego.AppConfig.DefaultString("RepositoryMode", "inmemory")
	if r == "postgres" {
		setupDB()
		controllers.InitPostgresRepo(orm.NewOrm())
	} else {
		controllers.InitInmemoryRepo()
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
