package main

import (
	"log"
	"net/http"

	"github.com/forderation/user-task/cmd"
	"github.com/forderation/user-task/internal/handler"
	"github.com/forderation/user-task/internal/middleware"
	"github.com/forderation/user-task/internal/repository/tasks"
	"github.com/forderation/user-task/internal/repository/users"
	"github.com/forderation/user-task/internal/usecase/task"
	"github.com/forderation/user-task/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	loadConfigFile()
	db := cmd.ProvideMysqlDB(viper.GetString("db_dsn"))
	handler := initRoute(db)
	address := viper.GetString("service_addr")
	srv := &http.Server{Addr: address, Handler: handler}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("listen: %s\n", err)
	}
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

	tasksRepository := tasks.NewRepository(db)
	usersRepository := users.NewRepository(db)

	taskUsecase := task.NewUsecase(task.NewUsecaseOptions{
		TaskRepository: tasksRepository,
	})
	userUsecase := user.NewUsecase(user.NewUsecaseOptions{
		UserRepository: usersRepository,
	})

	middleware := middleware.NewMiddleware(middleware.NewMiddlewareOptions{
		UserUsecase: userUsecase,
	})

	handler := handler.NewHandler(handler.NewHandlerOptions{
		UserUsecase: userUsecase,
		TaskUsecase: taskUsecase,
	})

	baseRoot.Use(middleware.CORS())

	userRoute := baseRoot.Group("/users")
	userRoute.POST("", handler.UserRegister)
	userRoute.POST("login", handler.UserLogin)
	userRoute.GET("", handler.UserGets)

	checkAuthorizationMiddleware := middleware.CheckUserAuthorization()

	userRoute.GET("/:id", checkAuthorizationMiddleware, handler.UserGet)
	userRoute.DELETE("/:id", checkAuthorizationMiddleware, handler.UserDelete)
	userRoute.PUT("/:id", checkAuthorizationMiddleware, handler.UserUpdate)

	tasksRoute := baseRoot.Group("/tasks", checkAuthorizationMiddleware)
	tasksRoute.POST("", handler.TaskPost)
	tasksRoute.GET("", handler.TaskGets)
	tasksRoute.GET("/:id", handler.TaskGet)
	tasksRoute.DELETE("/:id", handler.TaskDelete)
	tasksRoute.PUT("/:id", handler.TaskPut)

	return baseRoot
}
