package controllers

import (
	"github.com/entropyx/fiduchain/models"

	"github.com/kataras/iris/v12"
)

func (c *Controller) GetPhones(ctx iris.Context) {
	type message struct {
		Users []*models.User
	}
	//Getting list of users
	users := c.getUsers()
	msg := message{
		Users: users,
	}
	ctx.View("phones.html", msg)
}
