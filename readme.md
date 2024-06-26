# site-api
## 使用說明
1. 使用docker-compose安裝mysql, phpMyAdmin, redis, rabbitMQ
首先需要安裝`docker`
```
docker-compose up
```
之後要關閉請使用
```
docker-compose down
```
想查看資料庫可以到`http://localhost:8081`登入查看(phpMyAdmin)

新增.env，設定資料庫的環境變數
```
MYSQL_ROOT_PASSWORD = examplepassword
MYSQL_DATABASE = database
MYSQL_USER = user
MYSQL_PASSWORD = password
```

2. 開啟server
進入到目錄下使用`go run main.go`開啟server

3. 部署注意事項
    - 因為使用docker打包，所以要注意租用的server容量是否足夠，phpmyadmin是需要的，因為沒有撰寫post到DB的API
    - 之後前端部署上去後可能需要修改信任的代理IP，目前是使用`127.0.0.1`
    - 個人網站或許不需要redis，看需求做刪減
    - 連線資訊在docker和本地端上有所不同，本地IP要換成containerName

4. 打包說明
    `docker-compose up --build`
    `docker-compose build --no-cache`

    