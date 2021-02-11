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
	mysql.logger.Info().Msg("MySQL Server being started!")
	mysql.logger.Info().Msg("MySQL Server being started!")
	mysql.logger.Info().Msg("MySQL Server being started!")
	mysql.logger.Info().Msg("MySQL Server being started!")
	mysql.logger.Info().Msg("MySQL Server being started!")
	mysql.logger.Info().Msg("MySQL Server being started!")
	mysql.logger.Info().Msg("MySQL Server being started!")
	mysql.logger.Info().Msg("MySQL Server being started!")

	return nil
}
