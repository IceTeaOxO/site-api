package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "gorm.io/gorm"
)

// 大分類
type Category struct {
	// gorm.Model
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// 小分類
type Cat struct {
	// gorm.Model
	CatID      int    `json:"cat_id"`
	CatName    string `json:"cat_name"`
	CategoryID int    `json:"category_id"`
}

// post
type Post struct {
	// gorm.Model
	PostID       int       `json:"post_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Body         string    `json:"body"`
	FeatureImage string    `json:"feature_image"`
	Status       string    `json:"status"`
	Author       string    `json:"author"`
	CreatedAt    time.Time `json:"created_at"`
	CategoryID   int       `json:"category_id"`
	CatID        int       `json:"cat_id"`
}

// pf category
type PflCategory struct {
	// gorm.Model
	PflcatID   int    `json:"pflcat_id"`
	PflcatName string `json:"pflcat_name"`
}

// portfolio
type Portfolio struct {
	// gorm.Model
	PflID        int       `json:"pfl_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	FeatureImage string    `json:"feature_image"`
	URL          string    `json:"url"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	Tag          string    `json:"tag"`
	PflcatID     int       `json:"pflcat_id"`
}

// 定义一个结构体用于解析 JSON 数据
type PortfolioRequest struct {
	PflcatID string `json:"pflcat_id"`
}

type CatRequest struct {
	CatID string `json:"cat_id"`
}
type CategoryRequest struct {
	CategoryID string `json:"category_id"`
}
