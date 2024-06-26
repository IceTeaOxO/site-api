package mydb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func CreateTableSQL() {
	// // 指定要加載的.env文件的路徑，假設在特定資料夾中
	// envFilePath := ".env"

	// // 使用godotenv庫的Load函數加載環境變數
	// err := godotenv.Load(envFilePath)
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	// // 讀取環境變數的值
	// dbName := os.Getenv("MYSQL_DATABASE")
	// dbUser := os.Getenv("MYSQL_USER")
	// dbPassword := os.Getenv("MYSQL_PASSWORD")
	// dbHost := os.Getenv("MYSQL_HOST")

	// // 連接到 MySQL 資料庫
	// DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName))
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer DB.Close()

	db, err := InitializeDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 創建 category 資料表的 SQL 語句
	createCategoryTableSQL := `
	CREATE TABLE IF NOT EXISTS category (
		category_id INT AUTO_INCREMENT PRIMARY KEY,
		 category_name VARCHAR(255) NOT NULL
	);
	`
	// 創建 cat 資料表的 SQL 語句
	createCatTableSQL := `
	CREATE TABLE IF NOT EXISTS cat (
		cat_id INT AUTO_INCREMENT PRIMARY KEY,
		cat_name VARCHAR(255) NOT NULL,
		category_id INT,
		FOREIGN KEY (category_id) REFERENCES category(category_id)
	);
	`
	// 創建 posts 資料表的 SQL 語句
	createPostsTableSQL := `
	CREATE TABLE IF NOT EXISTS posts (
		post_id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description VARCHAR(255),
		body TEXT,
		feature_image VARCHAR(255),
		status ENUM('draft', 'published'),
		author VARCHAR(255),
		created_at DATETIME,
		category_id INT,
		cat_id INT,
		FOREIGN KEY (category_id) REFERENCES category(category_id),
		FOREIGN KEY (cat_id) REFERENCES cat(cat_id)
	);
	`
	// 創建 pfl_category 資料表的 SQL 語句
	createPflCategoryTableSQL := `
	CREATE TABLE IF NOT EXISTS pfl_category (
		pflcat_id INT AUTO_INCREMENT PRIMARY KEY,
		pflcat_name VARCHAR(255) NOT NULL
	);
	`
	// 創建 portfolio 資料表的 SQL 語句
	createPortfolioTableSQL := `
	CREATE TABLE IF NOT EXISTS portfolio (
		pfl_id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255),
		description VARCHAR(255),
		feature_image VARCHAR(255),
		url VARCHAR(255) NOT NULL,
		status ENUM('Activate', 'Deactivate'),
		created_at DATETIME,
		tag VARCHAR(255),
		pflcat_id INT,
		FOREIGN KEY (pflcat_id) REFERENCES pfl_category(pflcat_id)
	);
	`

	// 執行 SQL 命令以創建資料表
	_, err = db.Exec(createCategoryTableSQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(createCatTableSQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(createPostsTableSQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(createPflCategoryTableSQL)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(createPortfolioTableSQL)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("資料表創建成功！")
}

// 初始化資料庫連接
func InitializeDB() (*sql.DB, error) {
	var err error

	// 指定要加载的.env文件的路径，假设在特定文件夹中
	envFilePath := ".env"

	// 使用godotenv库的Load函数加载环境变量
	if loadErr := godotenv.Load(envFilePath); loadErr != nil {
		return nil, fmt.Errorf("error loading .env file: %v", loadErr)
	}

	// 读取环境变量的值
	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")

	// 連接到 MySQL 資料庫
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		panic(err.Error())
	}

	// 檢查連接是否成功
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection established")

	return DB, nil
}

// func InitializeGormDB() (*gorm.DB, error) {
// 	var err error

// 	// 指定要加载的.env文件的路径，假设在特定文件夹中
// 	envFilePath := ".env"

// 	// 使用godotenv库的Load函数加载环境变量
// 	if loadErr := godotenv.Load(envFilePath); loadErr != nil {
// 		return nil, fmt.Errorf("error loading .env file: %v", loadErr)
// 	}

// 	// 读取环境变量的值
// 	dbName := os.Getenv("MYSQL_DATABASE")
// 	dbUser := os.Getenv("MYSQL_USER")
// 	dbPassword := os.Getenv("MYSQL_PASSWORD")
// 	dbHost := os.Getenv("MYSQL_HOST")

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// 自动迁移模式，根据模型创建对应的表格
// 	db.AutoMigrate(&models.Category{})
// 	db.AutoMigrate(&models.Cat{})
// 	db.AutoMigrate(&models.Post{})
// 	db.AutoMigrate(&models.PflCategory{})
// 	db.AutoMigrate(&models.Portfolio{})

// 	return db, nil
// }
