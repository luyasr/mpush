package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/mpush/app/user"
	"github.com/luyasr/mpush/common/migrate"
	"github.com/luyasr/mpush/common/zerologger"
	"github.com/luyasr/mpush/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	var mode string
	if config.C.Server.Debug {
		mode = gin.DebugMode
	} else {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	r := gin.New()
	r.Use(zerologger.GinLogger(), zerologger.GinRecovery(true))
	u := user.NewHandler()
	v1 := r.Group("/api/v1")
	v1.Use()
	{
		v1.GET("ping", Ping)
		u.Registry(v1)
	}

	// automatically migrate databases
	migrate.AutoMigrate()

	start(r)
}

func start(r *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.C.Server.Port),
		Handler: r,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
