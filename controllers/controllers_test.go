package controllers

import (
	"fmt"
	"os"
	"testing"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/entropyx/fiduchain/models"
	. "github.com/smartystreets/goconvey/convey"
	//"gopkg.in/mgo.v2/bson"
)

func TestUsers(t *testing.T) {
	var prevUsers []*models.User
	mongoPath := fmt.Sprintf("%s:%s/%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB"))
	session, err := mgo.Dial(mongoPath)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	Convey("Given a new user", t, func() {
		user := &models.User{}
		Convey("When a random tel is set", func() {
			user.SetRandomTel()

			Convey("The tel should have length 10", func() {
				So(user.Tel, ShouldHaveLength, 10)
			})

			Convey("The tel should include LADA", func() {
				So(user.Tel, ShouldStartWith, "55")
			})
		})

		Convey("When a random verifying key is set", func() {
			user.SetRandomVerifyingKey()

			Convey("The verifying key should have length 10", func() {
				So(user.VerifyingKey, ShouldHaveLength, 10)
			})
		})

		Convey("When a random signing key is set", func() {
			user.SetRandomSigningKey()

			Convey("The signing key should have length 10", func() {
				So(user.SigningKey, ShouldHaveLength, 10)
			})
		})

		Convey("Given a controller", func() {
			ctr := New(session)
			Convey("When a previous list of users is retrieved", func() {
				prevUsers = ctr.getUsers()
			})

			Convey("When the user in inserted", func() {
				ctr.InsertUser(user)

				Convey("err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("The verifying key should not be empty", func() {
					So(user.VerifyingKey, ShouldNotBeBlank)
				})

				Convey("The signing key should not be empty", func() {
					So(user.SigningKey, ShouldNotBeBlank)
				})

				Convey("Given a user transaction", func() {
					transaction := &models.Transaction{
						Amount:       30000,
						VerifyingKey: user.VerifyingKey,
						SigningKey:   user.SigningKey,
						Timestamp:    int(time.Now().UnixNano()),
						UserTel:      user.Tel,
					}
					transaction.SetTimeLimit(3)

					Convey("When it is inserted into the blockchain", func() {
						err := ctr.InsertTransaction(transaction)

						Convey("err should be nil", func() {
							So(err, ShouldBeNil)
						})

						Convey("The transaction txid should not be empty", func() {
							So(transaction.TxId, ShouldNotBeBlank)
						})
					})
				})
			})

			Convey("When the list of users is retrieved", func() {
				users := ctr.getUsers()

				Convey("The list of uses should be greater than the insert", func() {
					So(len(users), ShouldBeGreaterThan, len(prevUsers))
				})
			})
		})
	})
}
