package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"path"
	"runtime"
	"strings"
	"sync"
)

var C = new(Config)

type Config struct {
	Server Server
	Mysql  Mysql
	Log    Log
}

type Server struct {
	Debug bool `json:"debug"`
	Port  int  `json:"port"`
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

func rootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	return root
}

func newConfig() {
	viper.AddConfigPath(rootPath())
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal errors config file: %s \n", err))
	}

	if err := viper.Unmarshal(C); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(C); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	viper.WatchConfig()
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
	var logMode logger.Interface
	if C.Server.Debug {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	if m.Conn == nil {
		m.Lock.Lock()
		defer m.Lock.Unlock()

		conn, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{
			Logger: logMode,
		})
		if err != nil {
			panic(err)
		}

		m.Conn = conn
	}

	return m.Conn
}

func init() {
	newConfig()
}
