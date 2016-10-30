package models

import "github.com/entropyx/fiduchain/utils"

type User struct {
	FullName     string `bson:"full_name,omitempty" json:"full_name"`
	Tel          string `bson:"tel,omitempty" json:"tel"`
	VerifyingKey string `bson:"verifying_key,omitempty" json:"verifying_key"`
	SigningKey   string `bson:"signing_key,omitempty" json:"signing_key"`
}

func (u *User) SetRandomName() {
	names := []string{
		"Pablo",
		"Juan",
		"Victoria",
		"Alejandro",
		"Daniel",
		"Pánfilo",
		"Brayan",
		"Estefani",
		"Kevin",
		"Yahaira",
		"Yadira",
		"Casiano",
		"Dilan",
		"Diosesano",
		"María",
		"Cristiano",
		"Gibrán",
		"Carolina",
	}
	lastnames := []string{
		"Pérez",
		"Gómez",
		"Mungía",
		"Puga",
		"Arvizu",
		"Montoya",
		"García",
		"Díaz",
		"Ramos",
		"Otazo",
	}
	name := names[utils.RandInt(0, len(names)-1)]
	lastname := lastnames[utils.RandInt(0, len(lastnames)-1)]
	u.FullName = name + " " + lastname
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
	u.SetRandomTel()
}
