package test

import (
	"testing"
	"time"
	"vec-calc-server/config"
	"vec-calc-server/lib"
	"vec-calc-server/model"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config := lib.LoadConfig[config.Config]()
	db = lib.NewDB(&config.DatabaseConfig, func(db *gorm.DB) error {
		return db.AutoMigrate(&model.User{})
	})
}

func TestUserRegister(t *testing.T) {
	user := model.User{
		Username: "testuser",
		Password: "testpassword",
	}
	var count int64
	if err := db.Where("username = ?", user.Username).Find(new([]model.User)).Count(&count).Error; count != 0 {
		t.Fatal("User already exists: ", err)
		return
	}
	user.Password = lib.HashedPassword(user.Password)
	if err := db.Create(&user).Error; err != nil {
		t.Fatal("Failed to create user: ", err)
		return
	}
}

func TestUserLogin(t *testing.T) {
	input, user := model.User{
		Username: "testuser",
		Password: "testpassword",
	}, new(model.User)
	if err := db.Where("username = ?", input.Username).Find(user).Error; err != nil {
		t.Fatal("Find user by email failed: ", err)
	}
	t.Log("input pass: ", input.Password, "user pass: ", user.Password)
	if err := lib.CheckPasswordHash(input.Password, user.Password); err != nil {
		t.Fatal("Incorrect password: ", err)
	}
	claim := lib.UserClaim{
		Name:   user.Username,
		Expire: time.Now().Add(time.Hour * 24),
	}
	token, err := lib.GenerateToken(&claim)
	if err != nil {
		t.Fatal("Failed to generate token: ", err)
	}
	t.Log("Token: ", token)
}
