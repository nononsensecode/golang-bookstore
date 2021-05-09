package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DatabaseConfig struct {
		EngineName string `yaml:"engine" env:"DB_ENGINE" env-default:"sqlite3" env-description:"Database engine/software name"`
		Filename string `yaml:"filename" env:"DB_FILENAME" env-default:"book.db" env-description:"Database file name for engines like sqlite, hsqldb"`
		Host string `yaml:"host" env:"DB_HOST" env-default:"localhost" env-description:"Domain name/IP of the computer hosting DB"`
		Port int32 `yaml:"port" env:"DB_PORT" env-default:"5432" env-description:"Port of the database server"`
		User string `yaml:"username" env:"DB_USER" env-default:"user" env-description:"Username of the database user"`
		Password string `yaml:"password" env:"DB_PASSWORD" env-default:"password" env-description:"Password of the database user"`
		DatabaseName string `yaml:"database" env:"DB_NAME" env-default:"books" env-description:"Database name for db's like postgres, mysql..etc"`
		DBOptions string `yaml:"options" env:"DB_OPTIONS" env-default:"" env-description:"Database options"`		
	} `yaml:"database"`
}

type args struct {
	ConfigFilePath string
	ConfigFilePrefix string
	Environment string
}

func (a *args) configFile() string {
	filename := fmt.Sprintf("%s-%s.yaml", a.ConfigFilePrefix, a.Environment)
	return filepath.Join(a.ConfigFilePath, filename)
}

func GetConfig() *Config {
	var cfg Config

	args := processArgs(&cfg)
	if err := cleanenv.ReadConfig(args.configFile(), &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return &cfg
}

func processArgs(cfg interface{}) *args {
	var argz args

	flag.StringVar(&argz.ConfigFilePath, "c", "./", "Path to configuration directory")
	flag.StringVar(&argz.ConfigFilePrefix, "p", "config", "Prefix of the configuration file")
	flag.StringVar(&argz.Environment, "e", "development", "Type of development environment such as production, testing..etc")

	fu := flag.CommandLine.Usage
	flag.CommandLine.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(flag.CommandLine.Output())
		fmt.Fprintln(flag.CommandLine.Output(), envHelp)
	}

	flag.Parse()
	return &argz
}