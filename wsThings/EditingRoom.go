package wsThings

import "net"

// WSConnection is wrapper over net.Conn to keep room id with the connection
type WSConnection struct {
	connection net.Conn
	send       chan []byte
}

// Message is a wrapper over simple ws message of byte array type, again to keep room id with it
type Message struct {
	data []byte
	room string
}

// EditingRoomManager is the central place to keep track of Rooms and users in rooms
type EditingRoomManager struct {
	Rooms           map[string]map[*WSConnection]bool
	BroadcastInRoom chan Message
	Register        chan Subscription
	Unregister      chan Subscription
}

func (h *EditingRoomManager) Init() {
	for {
		select {
		case s := <-h.Register:
			connections := h.Rooms[s.room]
			if connections == nil {
				connections = make(map[*WSConnection]bool)
				h.Rooms[s.room] = connections
			}
			h.Rooms[s.room][s.wsConn] = true
		case s := <-h.Unregister:
			connections := h.Rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.wsConn]; ok {
					delete(connections, s.wsConn)
					close(s.wsConn.send)
					if len(connections) == 0 {
						delete(h.Rooms, s.room)
					}
				}
			}
		case m := <-h.BroadcastInRoom:
			connections := h.Rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.Rooms, m.room)
					}
				}
			}
		}
	}
}

var EditingRoomManagerInstance = EditingRoomManager{
	BroadcastInRoom: make(chan Message),
	Register:        make(chan Subscription),
	Unregister:      make(chan Subscription),
	Rooms:           make(map[string]map[*WSConnection]bool),
}
