package mysql

import "time"

type MovieNightSession struct {
	id        int
	username  string
	sessionId string
	added     time.Time
}

func MovieNight(id int, username string, sessionId string, added time.Time) *MovieNightSession {
	return &MovieNightSession{
		id:        id,
		username:  username,
		sessionId: sessionId,
		added:     added,
	}
}

func (session *MovieNightSession) GetID() int {
	return session.id
}

func (session *MovieNightSession) GetUsername() string {
	return session.username
}

func (session *MovieNightSession) GetSessionId() string {
	return session.sessionId
}

func (session *MovieNightSession) GetAdded() time.Time {
	return session.added
}
