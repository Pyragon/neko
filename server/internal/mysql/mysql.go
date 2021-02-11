package mysql

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"n.eko.moe/neko/internal/types/config"
)

type MySQLHandler struct {
	logger       zerolog.Logger
	conf         *config.MySQL
	db           *sql.DB
	databaseName string
}

func New(conf *config.MySQL, database string) *MySQLHandler {
	logger := log.With().Str("module", "mysql").Logger()
	return &MySQLHandler{
		logger:       logger,
		conf:         conf,
		db:           nil,
		databaseName: database,
	}
}

func (mysql *MySQLHandler) Start() error {

	db, err := sql.Open("mysql", mysql.conf.DBUsername+":"+mysql.conf.DBPassword+"@tcp("+mysql.conf.DBHost+":"+strconv.Itoa(mysql.conf.DBPort)+")/"+mysql.databaseName+"?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	mysql.db = db

	return nil
}

func (mysql *MySQLHandler) GetAccount(id string) (*MovieNightSession, error) {

	var session *MovieNightSession

	if mysql.db == nil {
		return session, fmt.Errorf("DB not connected")
	}

	rows, err := mysql.db.Query("SELECT (id, username, session_id, added) FROM movie_night WHERE session_id=?", id)

	if err != nil {
		return session, fmt.Errorf("No user found")
	}

	for rows.Next() {
		err := rows.Scan(&session.id, &session.username, &session.sessionId, &session.added)

		if err != nil {
			return session, fmt.Errorf("Error scanning to struct")
		}
	}

	defer rows.Close()

	result := MovieNight(session.id, session.username, session.sessionId, session.added)

	return result, nil
}

func (mysql *MySQLHandler) GetPlayer(username string) (*PlayerDataType, error) {

	var player *PlayerDataType

	if mysql.db == nil {
		return player, fmt.Errorf("DB not connected")
	}

	rows, err := mysql.db.Query("", username)

	if err != nil {
		return player, fmt.Errorf("No user found")
	}

	for rows.Next() {

		err := rows.Scan(&player.id, &player.username, &player.rights)

		if err != nil {
			return player, fmt.Errorf("Error scanning to struct")
		}

	}

	defer rows.Close()

	result := PlayerData(player.id, player.username, player.rights)

	return result, nil

}

func (mysql *MySQLHandler) Shutdown() error {

	if mysql.db != nil {
		mysql.db.Close()
	}

	return nil
}