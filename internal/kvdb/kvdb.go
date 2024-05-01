package kvdb

type Player struct {
	SocketID   string `json:"socket_id"`
	Username   string `json:"username"`
	UserRank   string `json:"user_rank"`
	UserPoints int    `json:"user_points"`
	Room       string `json:"room"`
}
type Players struct {
	Players map[string]*Player
}

func NewPlayers() *Players {
	return &Players{
		Players: make(map[string]*Player),
	}
}

func (p *Players) AddPlayer(socketID string, username string, userRank string, userPoints int, room string) {
	p.Players[socketID] = &Player{
		SocketID:   socketID,
		Username:   username,
		UserRank:   userRank,
		UserPoints: userPoints,
		Room:       room,
	}
}

func (p *Players) RemovePlayer(socketID string) {
	delete(p.Players, socketID)
}

func (p *Players) GetPlayer(socketID string) *Player {
	return p.Players[socketID]
}

func (p *Players) GetAllPlayers() map[string]*Player {
	return p.Players
}

func (p *Players) UpdatePlayer(socketID string, username string, userRank string, userPoints int, room string) {
	p.Players[socketID].Username = username
	p.Players[socketID].UserRank = userRank
	p.Players[socketID].UserPoints = userPoints
	p.Players[socketID].Room = room
}

func NewPlayer() *Player {
	return &Player{
		SocketID:   "",
		Username:   "",
		UserRank:   "",
		UserPoints: 0,
		Room:       "",
	}
}

func (p *Player) SetPlayer(socketID string, username string, userRank string, userPoints int, room string) {
	p.SocketID = socketID
	p.Username = username
	p.UserRank = userRank
	p.UserPoints = userPoints
	p.Room = room
}
