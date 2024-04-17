package main

import (
	my_api "site-api/api"
	"site-api/mydb"
	"site-api/redis"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 設置信任的代理IP
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 取得所有blog大分類
	r.GET("/api/blog/category", my_api.GetCategory)
	// 根據條件取得blog小分類
	r.POST("/api/blog/cat", my_api.PostCat)
	// 取得所有blog文章
	r.GET("/api/blog/post", my_api.GetAllPost)
	// 根據條件取得blog文章
	r.POST("/api/blog/post/category", my_api.PostPostByCategory)
	// 根據條件取得blog文章
	r.POST("/api/blog/post/cat", my_api.PostPostByCat)
	// 根據id取得blog文章內容
	r.GET("/api/blog/post/:id", my_api.GetPostByID)
	// 取得所有portfolio大分類
	r.GET("/api/portfolio/category", my_api.GetPflCategory)
	// 取得所有portfolio
	r.GET("/api/portfolio", my_api.GetAllPortfolio)
	// 根據條件取得portfolio
	r.POST("/api/portfolio", my_api.PostPortfolio)
	// 根據id取得portfolio內容
	r.GET("/api/portfolio/:id", my_api.GetPortfolioByID)

	return r
}

func main() {
	mydb.CreateTableSQL()

	redis.InitializeRedis()

	r := setupRouter()

	r.Run(":8080")
}
