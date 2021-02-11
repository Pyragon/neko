package mysql

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"n.eko.moe/neko/internal/types/config"
)

type MySQLHandler struct {
	logger zerolog.Logger
	conf   *config.MySQL
	db     *sql.DB
}

func New(conf *config.MySQL) *MySQLHandler {
	logger := log.With().Str("module", "mysql").Logger()
	return &MySQLHandler{
		logger: logger,
		conf:   conf,
		db:     nil,
	}
}

func (mysql *MySQLHandler) Start() error {

	db, err := sql.Open("mysql", mysql.conf.DBUsername+":"+mysql.conf.DBPassword+"@tcp("+mysql.conf.DBHost+":"+strconv.Itoa(mysql.conf.DBPort)+")/"+mysql.conf.DBDatabase)

	if err != nil {
		panic(err.Error())
	}

	mysql.db = db

	mysql.logger.Info().Msg("Connected to MYSQL Server!")

	return nil
}

func (mysql *MySQLHandler) Shutdown() error {

	if mysql.db != nil {
		mysql.db.Close()
	}

	return nil
}
