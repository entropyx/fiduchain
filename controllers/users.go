package controllers

import (
	"github.com/entropyx/fiduchain/cli"
	"github.com/entropyx/fiduchain/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (c *Controller) insertUser(user *models.User) error {
	err := generateVerifyingKey(user)
	if err != nil {
		return err
	}
	col := c.Users()
	return col.Insert(user)
}

func (c *Controller) getUsers() []*models.User {
	var users []*models.User
	col := c.Users()
	col.Find(bson.M{}).All(&users)
	return users
}

func generateVerifyingKey(user *models.User) error {
	return cli.Get("keys", nil, user)
}

func (c *Controller) Users() *mgo.Collection {
	return c.DB("fiduchain").C("users")
}
