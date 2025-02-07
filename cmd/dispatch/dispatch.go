package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gucooing/hkrpg-go/dispatch"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "dispatch.json"
	err := dispatch.LoadConfig(confName)
	if err != nil {
		if err == dispatch.FileNotExist {
			p, _ := json.MarshalIndent(dispatch.DefaultConfig, "", "  ")
			cf, _ := os.Create("./conf/" + confName)
			cf.Write(p)
			cf.Close()
			fmt.Printf("找不到配置文件\n已生成默认配置文件 %s \n", confName)
			main()
		} else {
			panic(err)
		}
	}
	appid := alg.GetAppId()
	// 初始化日志
	logger.InitLogger("discord"+"["+appid+"]", strings.ToUpper(dispatch.GetConfig().LogLevel))
	logger.Info("hkrpg-go")

	cfg := dispatch.GetConfig()
	// 初始化
	newserver := dispatch.NewServer(cfg, appid)
	if newserver == nil {
		logger.Error("服务器初始化失败")
		return
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 启动SDK服务
	go func() {
		if err = newserver.Start(); err != nil {
			logger.Error("无法启动dispatch服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			logger.Info("dispatch服务正在关闭")
			if err = newserver.Shutdown(ctx); err != nil {
				logger.Error("无法正常关闭dispatch服务")
			}
			logger.Info("dispatch服务已停止")
			logger.CloseLogger()
			os.Exit(0)
		}
	}()
	select {}
}
