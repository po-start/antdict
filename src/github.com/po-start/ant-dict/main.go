package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"
	//"time"

	"github.com/po-start/ant-dict/core/config"
	"github.com/po-start/ant-dict/core/hack"
	"github.com/po-start/ant-dict/core/log"
	_ "github.com/po-start/ant-dict/core/db/mysql"
	"github.com/po-start/ant-dict/web"
)

var configFile *string = flag.String("config", "etc/config.yaml", "Item bank config file")
var logLevel *string = flag.String("log-level", "", "log level [debug|info|warn|error], default error")

func main() {
	fmt.Println("Start web server")
	fmt.Printf("PID: %d\n", os.Getpid()) // PID
	runtime.GOMAXPROCS(runtime.NumCPU()) // 分配cpu核数
	flag.Parse()                         // 守护进程
	fmt.Printf("Git commit:%s\n", hack.Version)
	fmt.Printf("Build time:%s\n", hack.Compile)
	// 加载配置文件
	conf, err := config.ParseConfigFile(*configFile)
	if err != nil {
		fmt.Printf("parse config file error:%v\n", err.Error())
		return
	}
	// 配置日志等级
	if *logLevel != "" {
		log.SetLogLevel(*logLevel)
	} else {
		log.SetLogLevel(conf.LOG_FILE)
	}
	// 配置日志文件
	if len(conf.LOG_PATH) != 0 {
		//dataTime := time.Now().Format("2006-01-02_15:04:05")
		LogFilePath := path.Join(conf.LOG_PATH, conf.LOG_FILE)
		err := log.New(LogFilePath, conf.LOG_SIZE, 1)
		if err != nil {
			fmt.Printf("new log file error:%v\n", err.Error())
			return
		}
		//  开启日志写入文件
		log.IsWriteFile(!conf.LOG_DEBUG)
	}

	// 实例化 web server
	webServer, err := web.NewWebServer(conf)
	if err != nil {
		fmt.Printf("Start web service error:%v\n", err.Error())
		return
	}

	// 信号处理
	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGPIPE,
	)

	go func() {
		for {
			sig := <-sc
			if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGQUIT {
				// 关闭各种连接
				log.Info("main", "main", "Got signal", 0, "signal", sig)
				fmt.Println("Shut down all kinds of connection")
				log.Close() // 关闭全部日志文件
			} else if sig == syscall.SIGPIPE {
				log.Info("main", "main", "Ignore broken pipe signal", 0)
			}
		}
	}()

	webServer.Run() // 启动web server
}
