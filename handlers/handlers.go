package handlers

import (
	"encoding/base64"
	"encoding/binary"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strings"
)

func LandingPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Posts",
	})
}

func EditorPage(c *gin.Context) {
	c.HTML(http.StatusOK, "editor.html", gin.H{
		"roomId": CreateRoom(c),
	})
}

func CreateRoom(c *gin.Context) string {
	roomId := createRoomId()
	return roomId
}

/**
 * Private functions. To be used only in this package
 */
func createRoomId() string {
	replacer := strings.NewReplacer("+", "-", "/", "-")
	randomNumber := rand.Int63()
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(randomNumber))
	encoded := base64.StdEncoding.EncodeToString([]byte(b))
	return replacer.Replace(encoded[:11])
}
