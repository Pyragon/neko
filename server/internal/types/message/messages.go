package message

import (
	"n.eko.moe/neko/internal/types"
)

type Message struct {
	Event string `json:"event"`
}

type Disconnect struct {
	Event   string `json:"event"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type SignalProvide struct {
	Event string   `json:"event"`
	ID    string   `json:"id"`
	SDP   string   `json:"sdp"`
	Lite  bool     `json:"lite"`
	ICE   []string `json:"ice"`
}

type SignalAnswer struct {
	Event string `json:"event"`
	SDP   string `json:"sdp"`
}

type MembersList struct {
	Event    string          `json:"event"`
	Memebers []*types.Member `json:"members"`
}

type Member struct {
	Event string `json:"event"`
	*types.Member
}
type MemberDisconnected struct {
	Event string `json:"event"`
	Name  string `json:"name"`
}

type Clipboard struct {
	Event string `json:"event"`
	Text  string `json:"text"`
}

type Keyboard struct {
	Event      string  `json:"event"`
	Layout     *string `json:"layout,omitempty"`
	CapsLock   *bool   `json:"capsLock,omitempty"`
	NumLock    *bool   `json:"numLock,omitempty"`
	ScrollLock *bool   `json:"scrollLock,omitempty"`
}

type Control struct {
	Event string `json:"event"`
	Name  string `json:"name"`
}

type ControlTarget struct {
	Event  string `json:"event"`
	Name   string `json:"name"`
	Target string `json:"target"`
}

type ChatReceive struct {
	Event   string `json:"event"`
	Content string `json:"content"`
}

type ChatRemove struct {
	Event string `json:"event"`
	ID    string `json:"id"`
}

type ChatAll struct {
	Event    string               `json:"event"`
	Messages []*types.ChatMessage `json:"messages"`
}

type ChatSend struct {
	Event       string             `json:"event"`
	Name        string             `json:"name"`
	ChatMessage *types.ChatMessage `json:"message"`
}

type EmoteReceive struct {
	Event string `json:"event"`
	Emote string `json:"emote"`
}

type EmoteSend struct {
	Event string `json:"event"`
	ID    string `json:"id"`
	Emote string `json:"emote"`
}

type Admin struct {
	Event string `json:"event"`
	Name  string `json:"name"`
}

type AdminError struct {
	Event string `json:"event"`
	Error string `json:"error"`
}

type AdminTarget struct {
	Event  string `json:"event"`
	Target string `json:"target"`
	Name   string `json:"name"`
}

type ScreenResolution struct {
	Event  string `json:"event"`
	ID     string `json:"id,omitempty"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Rate   int    `json:"rate"`
}

type ScreenConfigurations struct {
	Event          string                            `json:"event"`
	Configurations map[int]types.ScreenConfiguration `json:"configurations"`
}

type BroadcastStatus struct {
	Event    string `json:"event"`
	URL      string `json:"url"`
	IsActive bool   `json:"isActive"`
}

type BroadcastCreate struct {
	Event string `json:"event"`
	URL   string `json:"url"`
}
