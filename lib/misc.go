package lib

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// 为Gin router 添加CRUD
func APIBuilder(router gin.IRouter, handlers ...func(*gin.RouterGroup) *gin.RouterGroup) func(gin.IRouter, string) *gin.RouterGroup {
	return func(router gin.IRouter, path string) *gin.RouterGroup {
		group := router.Group(path)
		for _, handler := range handlers {
			group = handler(group)
		}
		return group
	}
}

// 生成token
func GenerateToken(userClaim *UserClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(token string) (*UserClaim, error) {
	userClaim := new(UserClaim)
	claim, err := jwt.ParseWithClaims(token, userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !claim.Valid {
		return nil, errors.New("token validation failed")
	}
	return userClaim, nil
}

// 验证token
func (uc *UserClaim) Valid() error {
	if time.Now().Unix() > uc.Expire.Unix() {
		return errors.New("token expired")
	}
	return nil
}

func HashedPassword(password string) string {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password: ", err)
	}
	return string(res)

}
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
