package controllers

import mgo "gopkg.in/mgo.v2"

type Controller struct {
	s *mgo.Session
}

func New(session *mgo.Session) *Controller {

	return &Controller{session}
}
