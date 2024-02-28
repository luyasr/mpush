package config

import (
	"fmt"
	"github.com/luyasr/gaia/config"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/stores/kafka"
	"github.com/luyasr/gaia/stores/mysql"
)

const (
	name = "config"
)

var Cfg = new(Config)

type Config struct {
	Http  Http          `json:"http"`
	Mysql *mysql.Config `json:"mysql"`
	Kafka *kafka.Config `json:"kafka"`
}

type Http struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Mode string `json:"mode"`
}

func (h *Http) Address() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func (c *Config) Init() error {
	return nil
}

func (c *Config) Name() string {
	return name
}

func init() {
	ioc.Container.Registry(ioc.ConfigNamespace, Cfg)

	cfg, err := config.New(config.LoadFile("config.yaml", Cfg))
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := cfg.Read(); err != nil {
		log.Fatal(err.Error())
	}
	if err := cfg.Watch(); err != nil {
		log.Fatal(err.Error())
	}
}
