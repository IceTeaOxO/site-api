# site-api
## 使用說明
1. 使用docker-compose安裝mysql, phpMyAdmin, redis, rabbitMQ
首先需要安裝`docker`
進入到`./deplotment`目錄運行以下指令
```
docker-compose up
```
之後要關閉請使用
```
docker-compose down
```
想查看資料庫可以到`http://localhost:8081`登入查看(phpMyAdmin)

在deployment底下新增.env，設定資料庫的環境變數
```
MYSQL_ROOT_PASSWORD = examplepassword
MYSQL_DATABASE = database
MYSQL_USER = user
MYSQL_PASSWORD = password
```

2. 開啟server
進入到目錄下使用`go run main.go`開啟server

