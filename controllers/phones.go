package controllers

import (
	"fmt"

	"github.com/kataras/iris"
)

func (c *Controller) GetPhones(ctx *iris.Context) {
	//Getting list of users
	users := c.getUsers()
	fmt.Println(users)
	ctx.ServeFile("templates/phones.html", false)
}
