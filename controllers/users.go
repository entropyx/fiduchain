package controllers

import (
	"github.com/entropyx/fiduchain/models"
	mgo "gopkg.in/mgo.v2"
)

func (c *Controller) insertUser(user *models.User) error {
	col := c.Users()
	col.Insert(user)
	return nil
}

func (c *Controller) Users() *mgo.Collection {
	return c.DB("fiduchain").C("users")
}
