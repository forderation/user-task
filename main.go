package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/forderation/user-task/cmd"
	"github.com/forderation/user-task/internal/repository/tasks"
	"github.com/forderation/user-task/internal/repository/users"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	loadConfigFile()

	// init driver
	db := cmd.ProvideMysqlDB(viper.GetString("db_dsn"))

	// init repository
	tasks.NewRepository(db)
	users.NewRepository(db)

	address := viper.GetString("service_addr")
	srv := &http.Server{Addr: address}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// gracefully shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown service ...")
	log.Println("service exiting")
}

func loadConfigFile() {
	filePath := fmt.Sprintf("config.toml")
	viper.SetConfigType("toml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("using config file:", viper.ConfigFileUsed())
}

func initRoute() *gin.Engine {
	baseRoot := gin.Default()
	baseRoot.Use(corsMiddleware())
	return baseRoot
}

func initMysqlDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Panic("error init mysql db: ", err.Error())
	}
	return db
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
