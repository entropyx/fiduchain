package controllers

import mgo "gopkg.in/mgo.v2"

type Controller struct {
	s *mgo.Session
}

func (c *Controller) DB(name string) *mgo.Database {
	return c.s.DB(name)
}

func New(session *mgo.Session) *Controller {
	return &Controller{session}
}
