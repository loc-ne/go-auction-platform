package websocket

type Room struct {
	id string
	clients map[*Client]bool
}

func NewRoom(id string) *Room {
	return &Room{
		id: id,
		clients: make(map[*Client]bool),
	}
}

func (r *Room) AddClient(c *Client) {
	r.clients[c] = true
}

func (r *Room) RemoveClient(c *Client) {
	delete(r.clients, c)
}

func (r *Room) Broadcast(msg Message) {
	for c := range r.clients {
		c.send <- msg
	}
}