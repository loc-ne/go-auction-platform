package websocket

import (
	"hash/maphash"
)

type Message struct {
	RoomID  string      `json:"roomId"`
	UserID  string      `json:"userId"`
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

var seed = maphash.MakeSeed()
func hash(id string) uint64 {
	var h maphash.Hash
	h.SetSeed(seed)
	h.WriteString(id)
	return h.Sum64()
}

type Hub struct {
	shards []*Shard
}

type Shard struct {
	rooms map[string]*Room

	register   chan *Client
	unregister chan *Client
	broadcast  chan Message
}

func NewHub() *Hub {
	h := &Hub{
		shards: make([]*Shard, 16),
	}
	for i := 0; i < 16; i++ {
		h.shards[i] = &Shard{
			rooms:      make(map[string]*Room),
			register:   make(chan *Client),
			unregister: make(chan *Client),
			broadcast:  make(chan Message),
		}
		go h.shards[i].Run() 
	}
	return h
}

func (h *Hub) GetShard(id string) *Shard {
	return h.shards[hash(id)%uint64(len(h.shards))]
}


func (h *Hub) Register(c *Client) {
	shard := h.GetShard(c.roomId)
	shard.register <- c
}

func (h *Hub) Unregister(c *Client) {
	shard := h.GetShard(c.roomId)
	shard.unregister <- c
}

func (h *Hub) Broadcast(msg Message) {
	shard := h.GetShard(msg.RoomID)
	shard.broadcast <- msg
}

func (s *Shard) Run() {
	for {
		select {
		case client := <-s.register:
			room, ok := s.rooms[client.roomId]
			if !ok {
				room = NewRoom(client.roomId)
				s.rooms[client.roomId] = room
			}
			room.AddClient(client)

		case client := <-s.unregister:
			if room, ok := s.rooms[client.roomId]; ok {
				room.RemoveClient(client)
				close(client.send) 
				
				if len(room.clients) == 0 {
					delete(s.rooms, client.roomId)
				}
			}
		case msg := <-s.broadcast:
			if room, ok := s.rooms[msg.RoomID]; ok {
				room.Broadcast(msg)
			}
		}
	}
}
