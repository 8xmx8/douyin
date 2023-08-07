package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/Godvictory/douyin/cmd/flags"
	"github.com/Godvictory/douyin/internal/conf"
	"github.com/Godvictory/douyin/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "前台启动服务",
	Long:  `Start the douyin server`,
	Run: func(cmd *cobra.Command, args []string) {
		initServer()
		if flags.Debug || flags.Dev {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		r := gin.New()
		server.Init(r)
		base := fmt.Sprintf("%s:%d", conf.Conf.Address, conf.Conf.Port)
		log.Infof("启动服务器 @ %s", base)
		srv := &http.Server{Addr: base, Handler: r}
		go func() {
			var err error
			if conf.Conf.Scheme.Https {
				err = srv.ListenAndServeTLS(conf.Conf.Scheme.CertFile, conf.Conf.Scheme.KeyFile)
			} else {
				err = srv.ListenAndServe()
			}
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("无法启动: %s", err.Error())
			}
		}()

		// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		log.Println("Server exiting")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
