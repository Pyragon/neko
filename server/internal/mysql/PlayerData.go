package mysql

type PlayerDataType struct {
	id       int
	username string
	rights   int
	banned   int
}

func PlayerData(id int, username string, rights int, banned int) PlayerDataType {
	return PlayerDataType{
		id:       id,
		username: username,
		rights:   rights,
		banned:   banned,
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

func (data *PlayerDataType) IsBanned() bool {
	return data.banned == 1
}
