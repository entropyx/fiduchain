package models

import (
	"github.com/entropyx/fiduchain/utils"
)

type User struct {
	Tel          string `bson:"tel,omitempty" json:"tel"`
	VerifyingKey string `bson:"verifying_key,omitempty" json:"verifying_key"`
	SigningKey   string `bson:"signing_key,omitempty" json:"signing_key"`
}

func (u *User) SetRandomTel() {
	lada := "55"
	str := utils.RandomIntString(8)
	u.Tel = lada + str
}

func (u *User) SetRandomVerifyingKey() {
	u.VerifyingKey = utils.RandomString(10)
}

func (u *User) SetRandomSigningKey() {
	u.SigningKey = utils.RandomString(10)
}

func (u *User) Fill() {
	u.SetRandomSigningKey()
	u.SetRandomVerifyingKey()
}
