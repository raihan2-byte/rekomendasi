package main

import (
	"log"
	"mygram/auth"
	"mygram/handler"
	"mygram/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/tugas2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connestion Error")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	// userInput := user.RegisterUserInput{}

	// userInput.Email = "raihan@gmail.com"
	// userInput.Username = "raihan"
	// userInput.Age = 10
	// userInput.Password = "inipassword"

	// userService.RegisterUser(userInput)

	db.AutoMigrate(&user.User{})

	// fmt.Println("Database Connection Success")

	router := gin.Default()
	api := router.Group("/users/v1")
	api.POST("/user", userHandler.RegisterUser)

	router.Run(":8080")
}
