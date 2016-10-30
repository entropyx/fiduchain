package controllers

import "github.com/kataras/iris"

func (c *Controller) GetPhones(ctx *iris.Context) {
	ctx.ServeFile("templates/phones.html", false)
}
