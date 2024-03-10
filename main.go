package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/forderation/user-task/cmd"
	"github.com/forderation/user-task/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	loadConfigFile()
	db := cmd.ProvideMysqlDB(viper.GetString("db_dsn"))
	handler := initRoute(db)
	address := viper.GetString("service_addr")
	srv := &http.Server{Addr: address, Handler: handler}
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
	filePath := "config.toml"
	viper.SetConfigType("toml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("using config file:", viper.ConfigFileUsed())
}

func initRoute(
	db *gorm.DB,
) *gin.Engine {
	baseRoot := gin.Default()

	middleware := middleware.NewMiddleware(middleware.NewMiddlewareOptions{})
	baseRoot.Use(middleware.CORS())

	return baseRoot
}

func initMysqlDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Panic("error init mysql db: ", err.Error())
	}
	return db
}
