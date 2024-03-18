package main

import (
	"embed"
	"net/http"
	"vec-calc-server/config"
	"vec-calc-server/lib"
	"vec-calc-server/model"

	"github.com/gin-gonic/gin"
)

//go:embed index.html
var fs embed.FS

func main() {
	config := lib.LoadConfig[config.Config]()
	router := gin.Default()
	calcRouter := router.Group("/api/calc")
	calcRouter.POST("/dot", HandleCalcDot)
	calcRouter.POST("/mul", HandleCalcMul)
	lib.AddStaticFS(router, fs)

	panic(router.Run(config.ServerConfig.Port))
}

// 非阻塞地进行请求处理
func HandleCalcDot(c *gin.Context) {
	var vectors []model.Vector
	if err := c.ShouldBindJSON(&vectors); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, func(v1 model.Vector, v2 model.Vector) float64 {
		if len(v1) != len(v2) {
			return 0
		}
		var result float64
		for i := range v1 {
			result += v1[i] * v2[i]
		}
		return result
	}(vectors[0], vectors[1]))
}
func HandleCalcMul(c *gin.Context) {
	var vectors []model.Vector
	if err := c.ShouldBindJSON(&vectors); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, func(v1 model.Vector, v2 model.Vector) model.Vector {
		if len(v1) != len(v2) {
			return nil
		}
		result := make(model.Vector, len(v1))
		for i := range v1 {
			result[i] = v1[i] * v2[i]
		}
		return result
	}(vectors[0], vectors[1]))
}
