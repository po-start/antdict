package web

import (
	"fmt"
//	"strconv"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/po-start/ant-dict/core/config"
	"github.com/po-start/ant-dict/core/log"
)

type WebServer struct {
	conf *config.Config
	*echo.Echo
}

func NewWebServer(conf *config.Config) (*WebServer, error) {
	webServer := new(WebServer)
	webServer.conf = conf
	webServer.Echo = echo.New()

	log.Info("web", "NewWebServer", "Web Server running", 0,
		"netProto",
		"http",
		"address",
		conf.ADDRESS, "prot", conf.PROT)

	return webServer, nil
}

func (self *WebServer) Run() error {
	self.RegisterMiddleware()
	self.RegisterURL()
	addr := fmt.Sprintf("%s:%d", self.conf.ADDRESS, self.conf.PROT)
	self.Server.Addr = addr
	self.Logger.Fatal(gracehttp.Serve(self.Server))
	return nil
}

func (self *WebServer) RegisterMiddleware() {
	/*self.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	    Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}",` +
	        `"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
	        `"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
	        `"bytes_out":${bytes_out}}` + "\n",
	    Output: log.GlobalSysLogger,
	}))*/
	self.Use(middleware.Logger())
	self.Use(middleware.Recover())
	//self.Use(middleware.BasicAuth(self.CheckAuth))
}

func (self *WebServer) CheckAuth(username, password string) bool {
	if username == self.conf.USER && password == self.conf.PASSWORD {
		return true
	}
	return false
}
