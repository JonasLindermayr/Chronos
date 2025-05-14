package main

import (
	"os"

	"github.com/JonasLindermayr/Chronos/handlers"
	"github.com/JonasLindermayr/Chronos/internal"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const Filepath = "data.db"

func main() {
	
	if _, err := os.Stat(Filepath); os.IsNotExist(err) {
		internal.Log("Database file not found. Please run migrate.go first.", internal.ERROR)
		panic(err)
	}
	var err error

	internal.DB, err = gorm.Open(sqlite.Open(Filepath), &gorm.Config{})
	if err != nil {
		internal.Log("Failed to connect to database", internal.ERROR)
		internal.Log("Error: " + err.Error(), internal.ERROR)
		panic(err)
	}
	_ = internal.DB
	internal.Log("Connected to database", internal.INFO)

	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	router.POST("/api/user", handlers.GetUserHandler)


	router.Run(":8072")
}