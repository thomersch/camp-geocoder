package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-cors"
)

type errorMsg struct {
	Message string `json:"msg"`
}

func main() {
	fmt.Println("Campsearch")

	pgConnStr := flag.String("pgConn", "postgres:///campcoder?sslmode=disable", "connection string for postgres")
	flag.Parse()
	err := bootDB(*pgConnStr)
	if err != nil {
		panic(err)
	}

	RunHTTP()
}

func RunHTTP() {
	router := gin.Default()
	router.Use(cors.Middleware(cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
	}))

	router.GET("/search", handleSearch)
	router.Run(":51009")
}

func handleSearch(c *gin.Context) {
	queryStr := c.Query("q")
	if len(queryStr) < 3 {
		c.JSON(400, errorMsg{"query has to have at least 3 chars"})
		return
	}
	res, err := Search(queryStr)
	if err != nil {
		c.JSON(500, errorMsg{fmt.Sprintf("internal error %s", err)})
		return
	}
	c.JSON(200, res)
}
