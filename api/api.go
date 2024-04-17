package my_api

import (
	"log"
	"net/http"
	"site-api/models"
	"site-api/mydb"
	"site-api/redis"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCategory(c *gin.Context) {
	// 構建 Redis 緩存鍵
	cacheKey := "categories"

	// 從 Redis 緩存中獲取結果
	cachedResult, err := redis.GetFromCache(cacheKey)
	if err == nil && cachedResult != "" {
		// 如果緩存存在，直接返回緩存結果
		c.JSON(http.StatusOK, cachedResult)
		return
	}

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

	// 將查詢結果轉換為 JSON 字串
	categoriesStr := redis.ConvertJsonToString(categories)

	// 將結果存入 Redis 緩存
	err = redis.SetToCache(cacheKey, categoriesStr, 5*time.Minute) // 設定 5 分鐘過期時間
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// POST:根據blog category條件取得blog小分類
func PostCat(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 定义一个变量用于接收解析后的 JSON 数据
	var request models.CategoryRequest

	// 解析请求中的 JSON 数据到结构体变量中
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// 查询数据库
	rows, err := db.Query("SELECT * FROM cat WHERE category_id = ?", request.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	// 解析查询结果
	var cats []models.Cat
	for rows.Next() {
		var cat models.Cat
		if err := rows.Scan(&cat.CatID, &cat.CatName, &cat.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		cats = append(cats, cat)
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"cats": cats})

}

func GetAllPost(c *gin.Context) {
	// 構建 Redis 緩存鍵
	cacheKey := "posts"

	// 從 Redis 緩存中獲取結果
	cachedResult, err := redis.GetFromCache(cacheKey)
	if err == nil && cachedResult != "" {
		// 如果緩存存在，直接返回緩存結果
		c.JSON(http.StatusOK, cachedResult)
		return
	}
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
	// 將查詢結果轉換為 JSON 字串
	postsStr := redis.ConvertJsonToString(posts)

	// 將結果存入 Redis 緩存
	err = redis.SetToCache(cacheKey, postsStr, 5*time.Minute) // 設定 5 分鐘過期時間
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"posts": posts})

}

// POST:根據blog cat條件取得blog文章
func PostPostByCat(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 定义一个变量用于接收解析后的 JSON 数据
	var request models.CatRequest

	// 解析请求中的 JSON 数据到结构体变量中
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// 查询数据库
	rows, err := db.Query("SELECT * FROM posts WHERE cat_id = ?", request.CatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	// 解析查询结果
	var posts []models.Post
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
			&post.CreatedAt,
			&post.CategoryID,
			&post.CatID,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		posts = append(posts, post)
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"posts": posts})

}

// POST:根據blog category條件取得blog文章
func PostPostByCategory(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 定义一个变量用于接收解析后的 JSON 数据
	var request models.CategoryRequest

	// 解析请求中的 JSON 数据到结构体变量中
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// 查询数据库
	rows, err := db.Query("SELECT * FROM posts WHERE category_id = ?", request.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	// 解析查询结果
	var posts []models.Post
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
			&post.CreatedAt,
			&post.CategoryID,
			&post.CatID,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		posts = append(posts, post)
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"posts": posts})

}

func GetPostByID(c *gin.Context) {
	// 構建 Redis 緩存鍵
	cacheKey := "post_" + c.Param("id")

	// 從 Redis 緩存中獲取結果
	cachedResult, err := redis.GetFromCache(cacheKey)
	if err == nil && cachedResult != "" {
		// 如果緩存存在，直接返回緩存結果
		c.JSON(http.StatusOK, cachedResult)
		return
	}
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 从 URL 参数中获取 portfolio ID
	id := c.Param("id")

	// 查询数据库
	rows, err := db.Query("SELECT * FROM posts WHERE post_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	// 解析查询结果
	var post models.Post
	if rows.Next() {
		if err := rows.Scan(
			&post.PostID,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.FeatureImage,
			&post.Status,
			&post.Author,
			&post.CreatedAt,
			&post.CategoryID,
			&post.CatID,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// 將查詢結果轉換為 JSON 字串
	postStr := redis.ConvertJsonToString(post)

	// 將結果存入 Redis 緩存
	err = redis.SetToCache(cacheKey, postStr, 5*time.Minute) // 設定 5 分鐘過期時間
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"post": post})

}

func GetPflCategory(c *gin.Context) {
	// 構建 Redis 緩存鍵
	cacheKey := "pfl_categories"

	// 從 Redis 緩存中獲取結果
	cachedResult, err := redis.GetFromCache(cacheKey)
	if err == nil && cachedResult != "" {
		// 如果緩存存在，直接返回緩存結果
		c.JSON(http.StatusOK, cachedResult)
		return
	}

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

	// 將查詢結果轉換為 JSON 字串
	pflCategoriesStr := redis.ConvertJsonToString(pflCategories)

	// 將結果存入 Redis 緩存
	err = redis.SetToCache(cacheKey, pflCategoriesStr, 5*time.Minute) // 設定 5 分鐘過期時間
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"pfl_categories": pflCategories})
}

func GetAllPortfolio(c *gin.Context) {
	// 構建 Redis 緩存鍵
	cacheKey := "portfolios"

	// 從 Redis 緩存中獲取結果
	cachedResult, err := redis.GetFromCache(cacheKey)
	if err == nil && cachedResult != "" {
		// 如果緩存存在，直接返回緩存結果
		c.JSON(http.StatusOK, cachedResult)
		return
	}
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

	// 將查詢結果轉換為 JSON 字串
	portfoliosStr := redis.ConvertJsonToString(portfolios)

	// 將結果存入 Redis 緩存
	err = redis.SetToCache(cacheKey, portfoliosStr, 5*time.Minute) // 設定 5 分鐘過期時間
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"portfolios": portfolios})
}

// POST:根據portfolio category條件取得portfolio
func PostPortfolio(c *gin.Context) {
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 定义一个变量用于接收解析后的 JSON 数据
	var request models.PortfolioRequest

	// 解析请求中的 JSON 数据到结构体变量中
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// 查询数据库
	rows, err := db.Query("SELECT * FROM portfolio WHERE pflcat_id = ?", request.PflcatID)
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

func GetPortfolioByID(c *gin.Context) {
	// 構建 Redis 緩存鍵
	cacheKey := "portfolio_" + c.Param("id")

	// 從 Redis 緩存中獲取結果
	cachedResult, err := redis.GetFromCache(cacheKey)
	if err == nil && cachedResult != "" {
		// 如果緩存存在，直接返回緩存結果
		c.JSON(http.StatusOK, cachedResult)
		return
	}
	db, err := mydb.InitializeDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing database"})
		return
	}
	defer db.Close()

	// 从 URL 参数中获取 portfolio ID
	id := c.Param("id")

	// 查询数据库
	rows, err := db.Query("SELECT * FROM portfolio WHERE pfl_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	// 解析查询结果
	var portfolio models.Portfolio
	if rows.Next() {
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
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	// 將查詢結果轉換為 JSON 字串
	portfolioStr := redis.ConvertJsonToString(portfolio)

	// 將結果存入 Redis 緩存
	err = redis.SetToCache(cacheKey, portfolioStr, 5*time.Minute) // 設定 5 分鐘過期時間
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"portfolio": portfolio})

}
