package types

type Member struct {
	ID    string `json:"id"`
	Name  string `json:"displayname"`
	Admin bool   `json:"admin"`
	Muted bool   `json:"muted"`
}

type Session interface {
	ID() string
	Name() string
	Admin() bool
	GetRights() int
	Muted() bool
	Connected() bool
	Member() *Member
	SetMuted(muted bool)
	GetLastMessage() int64
	SetLastMessage(time int64)
	SetConnected(connected bool) error
	SetSocket(socket WebSocket) error
	SetPeer(peer Peer) error
	Address() string
	Kick(message string) error
	Write(v interface{}) error
	Send(v interface{}) error
	SignalAnswer(sdp string) error
}

type SessionManager interface {
	New(id string, admin bool, rights int, username string, muted bool, socket WebSocket) Session
	HasHost() bool
	IsHost(name string) bool
	SetHost(name string) error
	GetHost() (Session, bool)
	ClearHost()
	Has(id string) bool
	Get(id string) (Session, bool)
	Members() []*Member
	Admins() []*Member
	Destroy(id string) error
	Clear() error
	Broadcast(v interface{}, exclude interface{}) error
	OnHost(listener func(id string))
	OnHostCleared(listener func(id string))
	OnDestroy(listener func(id string, session Session))
	OnCreated(listener func(id string, session Session))
	OnConnected(listener func(id string, session Session))
}
