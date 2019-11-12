package main

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/entropyx/fiduchain/controllers"
	"github.com/kataras/iris/v12"
)

// func main() {
// 	mongoPath := fmt.Sprintf("%s:%s/%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB"))
// 	session, err := mgo.Dial(mongoPath)
// 	defer session.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	ctrl := controllers.New(session)
// 	r := gin.Default()
// 	r.GET("/users/:tel/transactions", ctrl.GetUserTransactions)
// 	r.Run(":8081") // listen and server on 0.0.0.0:8080
// }
func main() {
	mongoPath := fmt.Sprintf("%s:%s/%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB"))
	session, err := mgo.Dial(mongoPath)
	defer session.Close()
	if err != nil {
		panic(err)
	}
	iris.RegisterOnInterrupt(session.Close)

	ctrl := controllers.New(session)
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html"))
	app.Get("/users/{tel}/transactions", ctrl.GetUserTransactions)
	app.Get("/users", ctrl.GetPhones)
	app.HandleDir("/public", "./public")
	app.Run(iris.Addr(":8081"))
}
