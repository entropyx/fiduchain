package controllers

import (
	"github.com/entropyx/fiduchain/models"
	"github.com/kataras/iris"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (c *Controller) GetUserTransactions(ctx *iris.Context) {
	transactions := c.getTransactionsByUserTel(ctx.Param("tel"))
	transactions = append(transactions, &models.Transaction{
		Amount:  1,
		UserTel: 234234})
	ctx.Render("transactions.html", transactions, iris.RenderOptions{"gzip": false, "charset": "UTF-8"})

}

func (c *Controller) InsertUserTransactions(context *gin.Context) {
	transactions := c.getTransactionsByUserTel(context.Param("tel"))
	context.JSON(200, transactions)
}

func (c *Controller) Transactions() *mgo.Collection {
	return c.DB("fiduchain").C("transactions")
}

func (c *Controller) getTransactionsByUserTel(tel string) []*models.Transaction {
	transactions := make([]*models.Transaction, 0)
	col := c.Transactions()
	col.Find(bson.M{"user_tel": tel})
	return transactions
}

func (c *Controller) insertUserTransactions() error {

	return nil
}

func (c *Controller) insertTransaction(transaction *models.Transaction) error {
	col := c.Transactions()
	return col.Insert(transaction)
}
