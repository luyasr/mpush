package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/luyasr/mpush/pkg/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"sync"
)

var (
	C    = new(Config)
	once sync.Once
)

type Config struct {
	Server Server
	Jwt    Jwt
	Mysql  Mysql
	Log    Log
}

type Server struct {
	Debug bool `json:"debug"`
	Port  int  `json:"port"`
}

type Jwt struct {
	Secret string `json:"secret"`
}

type Mysql struct {
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Database string   `json:"database"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Conn     *gorm.DB `json:"conn"`
	Lock     sync.Mutex
}

type Log struct {
	Dir string `json:"dir"`
}

func newConfig() {
	viper.AddConfigPath(utils.RootPath())
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal errors config file: %w", err))
	}

	if err := viper.Unmarshal(C); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err: %w", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(C); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err: %w", err))
		}
	})

	viper.WatchConfig()
}

func (m *Mysql) initDb() {
	var logMode logger.Interface
	if C.Server.Debug {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	once.Do(func() {
		conn, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{
			Logger: logMode,
		})
		if err != nil {
			panic(err)
		}

		m.Conn = conn
	})
}

func (m *Mysql) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)
}

func (m *Mysql) GetConn() *gorm.DB {
	return m.Conn
}

func init() {
	newConfig()
	C.Mysql.initDb()
}
