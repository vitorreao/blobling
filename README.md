# Blobling

Blobling is an open-source REST API for storing Binary Large OBjects (aka
BLOBs). It allows you to upload your BLOBs and organize them into folders.
Later, you can explore your stored BLOBs and download the ones you want.

## Run to run locally

Example of how you would go about running the API locally. First, start a
docker container for the Database. You can use postgres or mysql, but I will
use mysql in this example.

```sh
docker run --name blob-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret-blob -e MYSQL_DATABASE=blobling -e MYSQL_USER=weblob -e MYSQL_PASSWORD=weblob -d mysql
```

After the Database is up and running, you can start the application.

```sh
DATABASE_DRIVER=mysql DATABASE_URL="weblob:weblob@tcp(127.0.0.1:3306)/blobling?charset=utf8" PORT=8080 go run main.go
```
