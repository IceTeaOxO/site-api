package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 大分類
type Category struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// 小分類
type Cat struct {
	CatID      int    `json:"cat_id"`
	CatName    string `json:"cat_name"`
	CategoryID int    `json:"category_id"`
}

// post
type Post struct {
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
	PflcatID   int    `json:"pflcat_id"`
	PflcatName string `json:"pflcat_name"`
}

// portfolio
type Portfolio struct {
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
