package main

import (
	"log"
	"net/http"
	"site-api/models"
	"site-api/mydb"
	"site-api/redis"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCategory(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 查询数据库
	rows, err := db.Query("SELECT * FROM category")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	// 解析查询结果
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		categories = append(categories, category)
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func GetCat(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

}

func GetAllPost(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 查询数据库
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	// 解析查询结果
	var posts []models.Post
	print(rows)
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.PostID,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.FeatureImage,
			&post.Status,
			&post.Author,
			&post.CreatedAt, // 数据库中的时间字段可以直接转换为 time.Time 类型
			&post.CategoryID,
			&post.CatID,
		); err != nil {
			// 打印具体的错误信息以便调试
			log.Println("Error scanning rows:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		posts = append(posts, post)
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"posts": posts})

}

func GetPost(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

}

func GetPostByID(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

}

func GetPflCategory(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 查询数据库
	rows, err := db.Query("SELECT * FROM pfl_category")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	// 解析查询结果
	var pflCategories []models.PflCategory
	for rows.Next() {
		var pflCategory models.PflCategory
		if err := rows.Scan(
			&pflCategory.PflcatID,
			&pflCategory.PflcatName,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		pflCategories = append(pflCategories, pflCategory)
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"pfl_categories": pflCategories})
}

func GetAllPortfolio(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 查询数据库
	rows, err := db.Query("SELECT * FROM portfolio")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	// 解析查询结果
	var portfolios []models.Portfolio
	for rows.Next() {
		var portfolio models.Portfolio
		if err := rows.Scan(
			&portfolio.PflID,
			&portfolio.Title,
			&portfolio.Description,
			&portfolio.FeatureImage,
			&portfolio.URL,
			&portfolio.Status,
			&portfolio.CreatedAt,
			&portfolio.Tag,
			&portfolio.PflcatID,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		portfolios = append(portfolios, portfolio)
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"portfolios": portfolios})
}

func GetPortfolio(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

}

func GetPortfolioByID(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 取得所有blog大分類
	r.GET("/api/blog/category", GetCategory)
	// 根據條件取得blog小分類
	r.POST("/api/blog/cat", GetCat)
	// 取得所有blog文章
	r.GET("/api/blog/post", GetAllPost)
	// 根據條件取得blog文章
	r.POST("/api/blog/post", GetPost)
	// 根據id取得blog文章內容
	r.GET("/api/blog/post/:id", GetPostByID)
	// 取得所有portfolio大分類
	r.GET("/api/portfolio/category", GetPflCategory)
	// 取得所有portfolio
	r.GET("/api/portfolio", GetAllPortfolio)
	// 根據條件取得portfolio
	r.POST("/api/portfolio", GetPortfolio)
	// 根據id取得portfolio內容
	r.GET("/api/portfolio/:id", GetPortfolioByID)

	return r
}

func main() {
	mydb.CreateTableSQL()

	redis.InitializeRedis()

	r := setupRouter()

	r.Run(":8080")
}
