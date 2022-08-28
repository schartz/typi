package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pragati2112.github.io/db"
	"pragati2112.github.io/handlers"
	"pragati2112.github.io/wsThings"
)

func main() {

	go wsThings.EditingRoomManagerInstance.Init()

	// set up gin router with static file serving and HTML template rendering
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")
	router.Static("/static", "./static")

	// cutesy pingy pongy route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ws/:roomId", func(context *gin.Context) {
		wsThings.WsConnectionHandler(context)
	})

	router.GET("/", handlers.LandingPage)
	router.GET("/:roomId", handlers.EditorPage)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	db.DatabaseTest()
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
