package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type MySQL struct {
	DBUsername string
	DBPassword string
	DBDatabase string
	DBHost     string
	DBPort     int
}

func (MySQL) Init(cmd *cobra.Command) error {
	cmd.PersistentFlags().String("dbusername", "", "username for connecting to mysql")
	if err := viper.BindPFlag("dbusername", cmd.PersistentFlags().Lookup("dbusername")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("dbpassword", "", "password for connecting to mysql")
	if err := viper.BindPFlag("dbpassword", cmd.PersistentFlags().Lookup("dbpassword")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("dbdatabase", "", "database to connect to")
	if err := viper.BindPFlag("dbdatabase", cmd.PersistentFlags().Lookup("dbdatabase")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("dbhost", "", "host to connect to")
	if err := viper.BindPFlag("dbhost", cmd.PersistentFlags().Lookup("dbhost")); err != nil {
		return err
	}

	cmd.PersistentFlags().Int("dbport", 3306, "port to connect on")
	if err := viper.BindPFlag("dbport", cmd.PersistentFlags().Lookup("dbport")); err != nil {
		return err
	}

	return nil
}

func (s *MySQL) Set() {
	s.DBUsername = viper.GetString("dbusername")
	s.DBPassword = viper.GetString("dbpassword")
	s.DBDatabase = viper.GetString("dbdatabase")
	s.DBHost = viper.GetString("dbhost")
	s.DBPort = viper.GetInt("dbport")
}
