package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/kohTkd/experimental_learning/internal/util/environment"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	User                 string
	Password             string
	Net                  string
	Addr                 string
	Dbname               string
	AllowNativePasswords bool
	Params               struct {
		ParseTime string
		Charset   string
		Loc       string
	}
}

type Config struct {
	Database struct {
		Reader DatabaseConfig
		Writer DatabaseConfig
	}
	Server struct {
		Port int
	}
}

type ReadConfigOption struct {
	Env environment.EnvName
}

var Conf Config

func Read(opt ReadConfigOption) Config {
	row, err := os.ReadFile(filepath.Join(rootDir(), "config", "config.yml"))
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	expanded := []byte(os.ExpandEnv(string(row)))

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadConfig(bytes.NewBuffer(expanded)); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	c := &Conf

	if err := viper.UnmarshalKey(opt.Env.String(), &c); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return *c
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
