package lib

import (
	"embed"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// 通用配置
type ServerConfig struct {
	Port string `json:"port"`
	Mode string `json:"mode"`
}

// 配置管理
func LoadConfig[Config any]() *Config {
	if _, err := os.Stat("config.yaml"); err != nil {
		confTmpl := new(Config)
		data, _ := yaml.Marshal(confTmpl)
		os.WriteFile("config.yaml", []byte(data), 0644)
		log.Fatal(errors.New("config file not found, edit the config.yaml file and restart the server"))
	}
	if err := func() error {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		return viper.ReadInConfig()
	}(); err != nil {
		log.Fatal(errors.New("config file read failed"))
	}
	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		log.Fatal(errors.New("config file parse failed"))
	}
	return config
}

func AddStaticFS(router *gin.Engine, fs embed.FS) {
	router.NoRoute(gin.WrapH(http.FileServer(http.FS(fs))))
}
