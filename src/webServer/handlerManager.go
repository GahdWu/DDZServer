package webServer

import (
	"fmt"

	"moqikaka.com/goutil/debugUtil"
)

var (
	//Web业务处理器集合
	handlerMap = make(map[string]*Handler)
)

//注册处理器
func RegisteHandler(handler *Handler) {
	// 判断是否已经注册过，避免命名重复
	if _, exists := handlerMap[handler.Path]; exists {
		panic(fmt.Sprintf("%s已经被注册过，请重新命名", handler.Path))
	}
	handlerMap[handler.Path] = handler
	debugUtil.Println("AddHandler:", handler.Path)
}

//注销处理器
func UnRegisteHandler(path string) {
	delete(handlerMap, path)
}

//获取处理器
func GetHandler(path string) *Handler {
	if handler, exists := handlerMap[path]; exists {
		return handler
	}

	return nil
}
