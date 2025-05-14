package controller

import (
	"net/http"

	"github.com/JonasLindermayr/Chronos/internal"
	"github.com/JonasLindermayr/Chronos/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c *gin.Context){

	var authInput types.AuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		internal.Log("Failed to bind JSON", internal.ERROR)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var userFound types.User
	internal.DB.Where("username = ?", authInput.Username).First(&userFound)

	if userFound.ID == 0 {
		internal.Log("User not found", internal.ERROR)
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		internal.Log("Invalid password", internal.ERROR)
		internal.Log("Username: " + authInput.Username, internal.WARNING)
		internal.Log("IP: " + c.ClientIP(), internal.DEBUG)
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid password"})
		return
	}
	internal.Log("User logged in as: " + userFound.Username, internal.INFO)

	c.JSON(http.StatusOK, gin.H{
		"id": userFound.ID,
		"username": userFound.Username,
		"email": userFound.Email,
	})
}

func CreateUser(username, password, email string) {

	authInput := &types.AuthInput{
		Username: username,
		Password: password,
	}

	var userFound types.User
	internal.DB.Where("username = ?", authInput.Username).First(&userFound)
	
	if userFound.ID != 0 {
		internal.Log("User already exists", internal.ERROR)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		internal.Log("Failed to hash password", internal.ERROR)
	}

	user := &types.User{
		Username: authInput.Username,
		Password: string(passwordHash),
		Email: email,
	}

	internal.DB.Create(&user)
	internal.Log("User created", internal.INFO)

}

func CreateUserWithMigrate(username, password, email string) {

	authInput := &types.AuthInput{
		Username: username,
		Password: password,
	}

	var userFound types.User
	internal.DB.Where("username = ?", authInput.Username).First(&userFound)
	
	if userFound.ID != 0 {
		internal.LogMigrate("User already exists", internal.ERROR)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		internal.LogMigrate("Failed to hash password", internal.ERROR)
		return
	}

	user := &types.User{
		Username: authInput.Username,
		Password: string(passwordHash),
		Email: email,
	}

	internal.DB.Create(&user)
	internal.LogMigrate("User created", internal.INFO)

}