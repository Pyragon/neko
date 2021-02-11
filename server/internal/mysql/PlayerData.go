package mysql

type PlayerDataType struct {
	id       int
	username string
	rights   int
}

func PlayerData(id int, username string, rights int) *PlayerDataType {
	return &PlayerDataType{
		id:       id,
		username: username,
		rights:   rights,
	}
}

func (data *PlayerDataType) GetID() int {
	return data.id
}

func (data *PlayerDataType) GetUsername() string {
	return data.username
}

func (data *PlayerDataType) GetRights() int {
	return data.rights
}
