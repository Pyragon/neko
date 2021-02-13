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
	logger.Info().Msg("Username from NEW: " + conf.DBUsername)
	return &MySQLHandler{
		logger:       logger,
		conf:         conf,
		db:           nil,
		databaseName: database,
	}
}

func (mysql *MySQLHandler) Connect() *sql.DB {
	db, err := sql.Open("mysql", mysql.conf.DBUsername+":"+mysql.conf.DBPassword+"@tcp("+mysql.conf.DBHost+":"+strconv.Itoa(mysql.conf.DBPort)+")/"+mysql.databaseName+"?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func (mysql *MySQLHandler) GetAccount(id string) (MovieNightSession, error) {

	var session MovieNightSession

	db := mysql.Connect()

	rows, err := db.Query("SELECT id, username, session_id, expiry FROM sessions WHERE session_id=?", id)

	if err != nil {
		return session, fmt.Errorf("No user found: " + id + ", " + err.Error())
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

func (mysql *MySQLHandler) GetPlayer(username string) (PlayerDataType, error) {

	var player PlayerDataType

	db := mysql.Connect()

	rows, err := db.Query("SELECT id, username, rights, muted_from_movie_night, banned_from_movie_night FROM player_data WHERE username=?", username)

	if err != nil {
		return player, fmt.Errorf("No user found")
	}

	for rows.Next() {

		err := rows.Scan(&player.id, &player.username, &player.rights, &player.muted, &player.banned)

		if err != nil {
			return player, fmt.Errorf("Error scanning to struct")
		}

	}

	defer rows.Close()

	result := PlayerData(player.id, player.username, player.rights, player.muted, player.banned)

	return result, nil

}

func (mysql *MySQLHandler) Shutdown() error {

	if mysql.db != nil {
		mysql.db.Close()
	}

	return nil
}
