package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"vec-calc-server/model"

	"github.com/gin-gonic/gin"
)

// 非阻塞地进行请求处理
func HandleCalcDot(c *gin.Context) {
	var vectors [2]model.Vector
	if err := c.ShouldBindJSON(&vectors); err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, func(v1 model.Vector, v2 model.Vector) float64 {
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
	var vectors [2]model.Vector
	if err := c.ShouldBindJSON(&vectors); err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, func(v1 model.Vector, v2 model.Vector) model.Vector {
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

func BenchmarkAPI(b *testing.B) {
	// Prepare the request payloads
	vectors := func(numCases int) [][]model.Vector {
		testCases := make([][]model.Vector, numCases)
		for i := 0; i < numCases; i++ {
			vectorLength := rand.Intn(10000) + 1
			vectors := make([]model.Vector, 2)
			for j := 0; j < 2; j++ {
				vector := make(model.Vector, vectorLength)
				for k := 0; k < vectorLength; k++ {
					vector[k] = rand.Float64()
				}
				vectors[j] = vector
			}
			testCases[i] = vectors
		}
		return testCases
	}(64)

	payloads := make([][]byte, len(vectors))
	for i, v := range vectors {
		payload, _ := json.Marshal(v)
		payloads[i] = payload
	}

	router := gin.Default()
	apiRouter := router.Group("/api")
	apiRouter.POST("/calc/mul", HandleCalcMul)
	apiRouter.POST("/calc/dot", HandleCalcDot)

	// Perform the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j, payload := range payloads {
			b.Run(fmt.Sprintf("case-%d", j), func(b *testing.B) {
				req, _ := http.NewRequest("POST", "/api/calc/mul", bytes.NewBuffer(payload))
				req.Header.Set("Content-Type", "application/json")

				rr := httptest.NewRecorder()
				router.ServeHTTP(rr, req)
			})
			b.Run(fmt.Sprintf("case-%d", j), func(b *testing.B) {
				req, _ := http.NewRequest("POST", "/api/calc/dot", bytes.NewBuffer(payload))
				req.Header.Set("Content-Type", "application/json")

				rr := httptest.NewRecorder()
				router.ServeHTTP(rr, req)
			})
		}
	}
}
