package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const MAX_MUX = 10

var muxes = [MAX_MUX]sync.Mutex{}
var data = map[int]int{}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < MAX_MUX; i++ {
		muxes[i] = sync.Mutex{}
	}
	r := gin.New()
	r.Use(jsonLoggerMiddleware())
	r.GET("/p1", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to P1",
		})
	})
	r.GET("/p2", func(c *gin.Context) {

		e := rand.Intn(len(muxes))
		log.Println(e)
		muxes[e].Lock()
		time.Sleep(100 * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to P2",
		})
		muxes[e].Unlock()

	})
	r.GET("/p3", func(c *gin.Context) {

		e := rand.Intn(len(muxes))
		log.Println(e)
		muxes[e].Lock()
		time.Sleep(time.Duration(10^e) * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to P3",
		})
		muxes[e].Unlock()

	})
	r.GET("/p4", func(c *gin.Context) {
		randomTime := rand.Intn(10000)
		var values = make([]string, randomTime)
		for i := 0; i < randomTime; i++ {
			values[i] = RandStringRunes(100)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to P4: " + values[len(values)-1],
		})

	})
	r.Run("0.0.0.0:9000")
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func jsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("2006/01/02 - 15:04:05")
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}
