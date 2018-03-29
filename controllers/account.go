package controllers

import (
	"github.com/astaxie/beego"

	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type AccountController struct {
	beego.Controller
}

type Account struct {
	Id   int
	Name string
}

func (c *AccountController) Get() {

	host := beego.AppConfig.String("mysql_host")
	port := beego.AppConfig.String("mysql_port")
	user := beego.AppConfig.String("mysql_user")
	pass := beego.AppConfig.String("mysql_pass")
	database := beego.AppConfig.String("mysql_database")

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+database+"?charset=utf8")
	if err != nil {
		fmt.Printf("connect err")
	}

	rows, err1 := db.Query("select id, account_name from account limit 0,5")
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}

	defer rows.Close()

	//fmt.Println("")
	//cols, _ := rows.Columns()
	//for i := range cols {
	//	fmt.Print(cols[i])
	//	fmt.Print("\t")
	//}
	//fmt.Println("")

	var accounts []Account

	for rows.Next() {

		var id int
		var account_name string
		if err := rows.Scan(&id, &account_name); err == nil {
			//fmt.Print(id)
			//fmt.Print("\t")
			//fmt.Print(account_name)
			//fmt.Print("\t\r\n")
			accounts = append(accounts, Account{id, account_name})
		}
	}

	c.Data["accounts"] = accounts
	c.TplName = "account.tpl"
}
