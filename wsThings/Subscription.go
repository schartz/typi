package wsThings

import (
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// Subscription is a wrapper around the active ws connection from a user.
// It also keeps track of the user's joined room
type Subscription struct {
	wsConn *WSConnection
	room   string
}

func (s Subscription) WriteToSubscription() {
	defer func() {
		EditingRoomManagerInstance.Unregister <- s
		err := s.wsConn.connection.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		msg := <-s.wsConn.send
		err := wsutil.WriteServerMessage(s.wsConn.connection, ws.OpCode(ws.StateServerSide), msg)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}

func (s Subscription) ReadFromSubscription() {

	// when done, unregister from the EditingRoomManager
	// and close the connection
	defer func() {
		EditingRoomManagerInstance.Unregister <- s
		err := s.wsConn.connection.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		// read the incoming websocket message
		msg, _, err := wsutil.ReadClientData(s.wsConn.connection)
		if err != nil {
			fmt.Println(err)
			break
		}

		// wrap it into our room aware message wrapper
		m := Message{msg, s.room}

		// Sent this message to everyone present in the room
		EditingRoomManagerInstance.BroadcastInRoom <- m
	}
}
