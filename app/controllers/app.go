package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"my-app/app"
	"strings"
	"time"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Register(account, password string ) revel.Result {
	c.Validation.Email(account).Message("邮箱格式不正确")
	c.Validation.MinSize(password,6).Message("密码太短")
	c.Validation.MaxSize(password,16).Message("密码太长")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	queryStr := fmt.Sprintf("INSERT INTO account (account, password, register_date) values ('%v', '%v', '%v')",
		account, password, strings.Split( time.Now().Format(time.RFC3339), "T")[0])
	_, err := app.DB.Exec(queryStr)
	if err != nil {
		c.Log.Errorf("insert to db error :%v\n",err)
		c.Flash.Error("register error, 服务器错误")
		return c.Redirect(App.Index)
	}
	return c.Render(account)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("your name is request")
	c.Validation.MinSize(myName,3).Message("your name is less than 3")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}