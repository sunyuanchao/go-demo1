package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sunyd/go-demo1/common"
)

func main() {

	db := common.InitSqlDb()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	r.Run() // listen a
}
