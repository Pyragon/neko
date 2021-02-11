package mysql

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"n.eko.moe/neko/internal/types/config"
)

type MySQLHandler struct {
	logger zerolog.Logger
	conf   *config.MySQL
}

func New(conf *config.MySQL) *MySQLHandler {
	logger := log.With().Str("module", "mysql").Logger()
	return &MySQLHandler{
		logger: logger,
		conf:   conf,
	}
}

func (mysql *MySQLHandler) Start() error {

	mysql.logger.Info().Msg("Username: " + mysql.conf.DBUsername)
	mysql.logger.Info().Msg("Password: " + mysql.conf.DBPassword)
	mysql.logger.Info().Msg("Database: " + mysql.conf.DBDatabase)

	return nil
}

func (mysql *MySQLHandler) Shutdown() error {
	return nil
}
