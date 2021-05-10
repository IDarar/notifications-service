package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Environment string
		Mongo       MongoConfig
		GRPC        GRPCConfig
	}
	MongoConfig struct {
		URI      string
		User     string
		Password string
		Name     string `mapstructure:"databaseName"`
	}
	GRPCConfig struct {
		Port             string `mapstructure:"port"`
		ServerCertFile   string `mapstructure:"servercertfile"`
		ServerKeyFile    string `mapstructure:"serverkeyfile"`
		ClientCACertFile string `mapstructure:"clientcacertfile"`
		ClientKeyFile    string `mapstructure:"clinetkeyfile"`
		ClientCertFile   string `mapstructure:"clinetcertfile"`
	}
)

func Init(path string) (*Config, error) {

	if err := parseEnv(); err != nil {
		return nil, err
	}

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)
	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("mongo", &cfg.Mongo); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("grpc", &cfg.GRPC); err != nil {
		return err
	}

	return nil
}
func setFromEnv(cfg *Config) {
	cfg.Mongo.URI = viper.GetString("uri")
	cfg.Mongo.User = viper.GetString("user")
	cfg.Mongo.Password = viper.GetString("pass")
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // config file name

	return viper.ReadInConfig()
}

func parseEnv() error {
	return parseMongoEnvVariables()
}

//TODO env variables should be set in docker container
func parseMongoEnvVariables() error {

	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	os.Setenv("MONGO_PASS", "rootpassword")

	os.Setenv("MONGO_USER", "root")

	viper.SetEnvPrefix("mongo")
	if err := viper.BindEnv("uri"); err != nil {
		return err
	}

	if err := viper.BindEnv("user"); err != nil {
		return err
	}

	return viper.BindEnv("pass")
}
