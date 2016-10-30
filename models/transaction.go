package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	SId       string        `bson:"sid,omitempty" json:"sid"`
	Amount    int           `bson:"amount,omitempty" json:"amount"`
	Timestamp int           `bson:"timestamp,omitempty" json:"timestamp"`
	TimeLimit int           `bson:"time_limit,omitempty" json:"time_limit"`
	UserTel   string        `bson:"user_tel,omitempty" json:"user_tel"`
}

func (t *Transaction) setTimeLimit(days int) {
	timestamp := int64(t.Timestamp)
	timeLimit := int(time.Unix(0, timestamp).AddDate(0, 0, days).UnixNano())
	t.TimeLimit = timeLimit
}

func (t *Transaction) IsLate() bool {
	return t.TimeLimit < t.Timestamp
}
