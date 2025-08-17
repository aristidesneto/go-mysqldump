package config

import (
	"errors"
	"log/slog"
	"os"
	"strings"

	"github.com/go-playground/validator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	AWS struct {
		Bucket string `mapstructure:"bucket" validate:"required"`
	} `mapstructure:"aws"`

	Log struct {
		Path  string `mapstructure:"path"`
		Level string `mapstructure:"level" validate:"oneof=debug info warn error fatal"`
	} `mapstructure:"log"`

	Storage struct {
		Directory string `mapstructure:"directory"`
	} `mapstructure:"storage"`

	Compress struct {
		Type string `mapstructure:"type" validate:"required,oneof=bzip2 zstd"`
	} `mapstructure:"compress"`

	Databases []struct {
		Name         string   `mapstructure:"name"`
		Charset      string   `mapstructure:"charset"`
		IgnoreTables []string `mapstructure:"ignore_tables"`
	} `mapstructure:"databases"`

	Mysql struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"mysql"`
}

var cfg Config

func GetConfig() Config {
	return cfg
}

func LoadConfig(cmd *cobra.Command, cfgFile string) error {
	viper.SetEnvPrefix("GO_DUMP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil
		}
	}

	// Bind das flags com ENVs
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil
	}

	InitLogger(viper.GetString("log.level"))

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		slog.Error("Erro validação:", "error", err)
		os.Exit(1)
	}

	slog.Debug("Using config file", "filename", viper.ConfigFileUsed())

	return nil
}
