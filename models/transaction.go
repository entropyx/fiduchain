package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	Id                 bson.ObjectId `bson:"_id,omitempty" json:"id"`
	TxId               string        `bson:"txid,omitempty" json:"txid"`
	Amount             int           `bson:"amount,omitempty" json:"amount"`
	AmountWithoutCents float64       `bson:"amount_without_cents,omitempty" json:"amount_without_cents"`
	Timestamp          int           `bson:"timestamp,omitempty" json:"timestamp"`
	TimeLimit          int           `bson:"time_limit,omitempty" json:"timelimit"`
	Late               bool          `bson:"late,omitempty" json:"late"`
	VerifyingKey       string        `bson:"-" json:"verifying_key"`
	SigningKey         string        `bson:"-" json:"signing_key"`
	UserTel            string        `bson:"user_tel,omitempty" json:"user_tel"`
}

func (t *Transaction) SetTimeLimit(days int) {
	timestamp := int64(t.Timestamp)
	timeLimit := int(time.Unix(0, timestamp).AddDate(0, 0, days).UnixNano())
	t.TimeLimit = timeLimit
}

func (t *Transaction) SetLate() {
	t.Late = t.IsLate()
}

func (t *Transaction) SetAmountWithoutCents() {
	t.AmountWithoutCents = float64(t.Amount) / 100
}

func (t *Transaction) IsLate() bool {
	return t.TimeLimit < t.Timestamp
}
