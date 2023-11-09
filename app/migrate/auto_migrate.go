package migrate

import (
	"github.com/luyasr/mpush/app/message"
	"github.com/luyasr/mpush/app/user"
	"github.com/luyasr/mpush/config"
)

func AutoMigrate() {
	db := config.C.Mysql.GetConn()
	err := db.AutoMigrate(&user.User{}, &message.Message{})
	if err != nil {
		panic(err)
	}
}
