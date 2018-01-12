package web

import (
	"github.com/po-start/ant-dict/application/controller"
)

func (self *WebServer) RegisterURL() {
	dictController := &controller.DictController{}
	self.GET("/search", dictController.Search)                      // 搜索
}
