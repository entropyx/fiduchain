package main

import (
	"fmt"
	"os"
	"time"

	"github.com/entropyx/fiduchain/controllers"
	"github.com/entropyx/fiduchain/models"
	"github.com/entropyx/fiduchain/utils"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	mongoPath := fmt.Sprintf("%s:%s/%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB"))
	session, err := mgo.Dial(mongoPath)
	defer session.Close()
	if err != nil {
		panic(err)
	}
	ctrl := controllers.New(session)
	for i := 0; i < 20; i++ {
		user := &models.User{}
		err := ctrl.InsertUser(user)
		if err != nil {
			panic(err)
		}
		for j := 0; j < 20; j++ {
			transaction := &models.Transaction{
				Amount:       utils.RandInt(-300000, 300000),
				VerifyingKey: user.VerifyingKey,
				SigningKey:   user.SigningKey,
				Timestamp:    int(time.Now().UnixNano()),
				UserTel:      user.Tel,
			}
			transaction.SetTimeLimit(utils.RandInt(-3, 8))
			err := ctrl.InsertTransaction(transaction)
			if err != nil {
				panic(err)
			}
		}
	}
}
