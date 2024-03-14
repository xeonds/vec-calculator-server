package lib

import "time"

type UserClaim struct {
	Permission int       `json:"permission"`
	Name       string    `json:"name"`
	Expire     time.Time `json:"expire"` // token过期时间
}

// 通用配置
type ServerConfig struct {
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Type     string `json:"type"` // 数据库类型
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"` // 数据库名
	Migrate  bool   `json:"migrate"`
}

var jwtSecret = "secret"
