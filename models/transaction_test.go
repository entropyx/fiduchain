package models

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	//"gopkg.in/mgo.v2/bson"
)

func TestTransaction(t *testing.T) {
	Convey("Given a transaction", t, func() {
		transaction := &Transaction{
			Amount:    10000,
			Timestamp: int(time.Now().UnixNano()),
		}

		Convey("When the date limit was 3 days ago", func() {
			transaction.setTimeLimit(-3)

			Convey("The time limit should be less than timestamp", func() {
				So(transaction.TimeLimit, ShouldBeLessThan, transaction.Timestamp)
			})

			Convey("The transaction should be late", func() {
				So(transaction.IsLate(), ShouldBeTrue)
			})
		})

		Convey("When the date limit is in the future", func() {
			transaction.setTimeLimit(3)

			Convey("The time limit should be greater than timestamp", func() {
				So(transaction.TimeLimit, ShouldBeGreaterThan, transaction.Timestamp)
			})

			Convey("The transaction should NOT be late", func() {
				So(transaction.IsLate(), ShouldBeFalse)
			})
		})
	})
}
