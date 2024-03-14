package lib

import (
	"embed"
	"log"
	"net/http"
	"time"

	"vec-calc-server/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddStaticFS(router *gin.Engine, fs embed.FS) {
	router.NoRoute(gin.WrapH(http.FileServer(http.FS(fs))))
}
func AddLoginAPI(router gin.IRouter, path string, db *gorm.DB) *gin.RouterGroup {
	return APIBuilder(router, func(group *gin.RouterGroup) *gin.RouterGroup {
		group.POST("/login", handleLogin(db))
		group.POST("/register", handleRegister(db))
		return group
	})(router, path)
}

// login service
func handleLogin(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		input, user := new(model.User), new(model.User)
		if err := c.ShouldBindJSON(input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Where("username = ?", input.Username).Find(user).Error; err != nil {
			log.Println("Find user by username failed: ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		if err := CheckPasswordHash(input.Password, user.Password); err != nil {
			log.Println("Incorrect password: ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		token, err := GenerateToken(&UserClaim{Name: user.Username, Expire: time.Now().Add(time.Hour * 24)})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
func handleRegister(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var user model.User
		var count int64
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Where("username = ?", user.Username).Find(new(model.User)).Count(&count).Error; count > 0 {
			log.Println("User already exists: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
			return
		}
		user.Password = HashedPassword(user.Password)
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
	}
}
