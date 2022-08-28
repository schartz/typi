package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"net/http"
	"time"
)

func nameConn(conn net.Conn) string {
	return conn.LocalAddr().String() + " > " + conn.RemoteAddr().String()
}

func handler(c *gin.Context) {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}()
}

func main() {
	router := gin.Default()

	router.GET("/unique-id", func(c *gin.Context) {
		fmt.Println(time.Now().UnixNano())
		c.JSON(http.StatusOK, gin.H{
			"roomId": time.Now().UnixNano(),
		})
	})

	router.Static("/static", "./static")

	router.GET("/ws", handler)

	router.LoadHTMLGlob("./templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Posts",
		})
	})

	router.GET("/:roomId", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editor.html", gin.H{
			"title": "Posts",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
