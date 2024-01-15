package main

import (
	"fmt"
	"hexa-gorm/core/handlers"
	_ "hexa-gorm/core/models"
	"hexa-gorm/core/repository"
	"hexa-gorm/core/service"
	"hexa-gorm/packages/logs"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
  // init config
  initConfig()

  // init timezone
  initTimezone()

  // init database
  db := initDatabase()

  // init repository
  userRepository := repository.NewUserRepository(db)
  userServices := service.NewUserService(userRepository)
  userHandler := handlers.NewUserHandler(userServices)

  // init router
  r := gin.Default()
  userRoute := r.Group("/users")
  {
    userRoute.GET("", userHandler.GetUsers)
  }

  // run server

  logs.Info("Server is running..."+ viper.GetString("app.port"))
  r.Run(fmt.Sprintf(":%s", viper.GetString("app.port")))
}

func initConfig() {
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")
  viper.AutomaticEnv()
  viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

  if err := viper.ReadInConfig(); err != nil {
    panic(err)
  }
}


func initTimezone() {
  // set timezone
  ict, err := time.LoadLocation("Asia/Vientiane")
  if err != nil {
    panic(err)
  }
  
  time.Local = ict
}
func initDatabase () *gorm.DB {
  // Connect to a Postgres database
  dsn := fmt.Sprintf("host=%v user=%s password=%s dbname=%v port=%v sslmode=%v TimeZone=%v",
    viper.GetString("db.host"),
    viper.GetString("db.user"),
    viper.GetString("db.password"),
    viper.GetString("db.name"),
    viper.GetString("db.port"),
    viper.GetString("db.sslmode"),
    viper.GetString("db.tz"),
    )
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  // Migrate the schema
  // db.AutoMigrate(
  //   &models.User{},
  // )

  return db
}

