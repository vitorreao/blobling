package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vitorreao/blobling/user"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbDriver := getDatabaseDriver()
	dbURL := getDatabaseURL()
	db, err := getDatabase(dbDriver, dbURL)

	if err != nil {
		log.Fatalf("Could not connect to database: %s", err.Error())
	}

	router := gin.New()
	router.Use(gin.Logger())
	api := router.Group("/api")
	api.Use(user.ValidateSession)
	user.AddRoutes(api.Group("/users"), db)

	port := getPort()
	router.Run(":" + port)
}

func getDatabaseDriver() string {
	var dbDriver string
	flag.StringVar(&dbDriver, "db_driver", os.Getenv("DATABASE_DRIVER"), "Database driver. Accept postgres or mysql")

	if dbDriver == "" {
		log.Fatal("$DATABASE_DRIVER must be set or --db_driver flag must be provided")
	}

	return strings.TrimSpace(strings.ToLower(dbDriver))
}

func getDatabaseURL() string {
	var dbURL string
	flag.StringVar(&dbURL, "db_url", os.Getenv("DATABASE_URL"), "Database URL")

	if dbURL == "" {
		log.Fatal("$DATABASE_URL must be set or --db_url flag must be provided")
	}

	return strings.TrimSpace(strings.ToLower(dbURL))
}

func getDatabase(driver string, url string) (*gorm.DB, error) {
	var dialector gorm.Dialector
	if driver == "postgres" {
		dialector = postgres.Open(url)
	} else if driver == "mysql" {
		dialector = mysql.Open(url)
	} else {
		return nil, fmt.Errorf("Invalid database driver %s", driver)
	}
	return gorm.Open(dialector, &gorm.Config{})
}

func getPort() string {
	var port string
	portEnv := os.Getenv("PORT")
	flag.StringVar(&port, "port", portEnv, "Port on which to listen")
	flag.Parse()

	if port == "" {
		log.Fatal("$PORT must be set or --port flag must be provided")
	}

	return port
}
