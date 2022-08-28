package wsThings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
)

func WsConnectionHandler(c *gin.Context) {
	roomId := c.Param("roomId")
	fmt.Println(roomId)
	fmt.Println("***********************")

	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		fmt.Println(err)
	}

	// create wrapped connection and subscription instances
	connection := &WSConnection{conn, make(chan []byte, 1024)}
	subscription := Subscription{connection, roomId}

	// add this subscription into EditingRoomManager
	EditingRoomManagerInstance.Register <- subscription

	// start reader and writer go routines on the subscription
	go subscription.WriteToSubscription()
	go subscription.ReadFromSubscription()

	//go func() {
	//	defer conn.Close()
	//
	//	for {
	//		msg, op, err := wsutil.ReadClientData(conn)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		err = wsutil.WriteServerMessage(conn, op, msg)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//
	//	}
	//}()
}
