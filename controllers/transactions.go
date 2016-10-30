package controllers

import (
	"fmt"

	"github.com/entropyx/fiduchain/cli"
	"github.com/entropyx/fiduchain/models"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (c *Controller) GetUserTransactions(ctx *iris.Context) {
	var debt int
	var payment int
	transactions := c.getTransactionsByUserTel(ctx.Param("tel"))
	fmt.Println("transactions", transactions)
	type message struct {
		Phone        string
		Debt         float64
		Payment      float64
		Balance      string
		Transactions []*models.Transaction
	}
	for i := 0; i < len(transactions); i++ {
		if transactions[i].Amount < 0 {
			debt += (transactions[i].Amount * -1)
		} else {
			payment += transactions[i].Amount
		}
	}

	msg := message{
		Phone:        ctx.Param("tel"),
		Debt:         float64(debt) / 100,
		Payment:      float64(payment) / 100,
		Balance:      fmt.Sprintf("%.2f", ((float64(payment) / 100) - float64(debt)/100)),
		Transactions: transactions,
	}
	ctx.Render("transactions.html", msg, iris.RenderOptions{"gzip": false, "charset": "UTF-8"})

}

func (c *Controller) InsertUserTransaction(context *gin.Context) {
	transactions := c.getTransactionsByUserTel(context.Param("tel"))
	context.JSON(200, transactions)
}

func (c *Controller) Transactions() *mgo.Collection {
	return c.DB("fiduchain").C("transactions")
}

func (c *Controller) getTransactionsByUserTel(tel string) []*models.Transaction {
	var transactions []*models.Transaction
	col := c.Transactions()
	col.Find(bson.M{"user_tel": tel}).All(&transactions)
	return transactions
}

func (c *Controller) InsertTransaction(transaction *models.Transaction) error {
	err := cli.Post("insert", transaction, transaction)
	if err != nil {
		return err
	}
	transaction.SetAmountWithoutCents()
	transaction.SetLate()
	col := c.Transactions()
	return col.Insert(transaction)
}
