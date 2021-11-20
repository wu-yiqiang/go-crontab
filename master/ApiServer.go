package master

import (
	"net"
	"net/http"
	"time"
	"strconv"
)

type ApiServer struct {
	httpServer *http.Server
}

func handlerJobSave(w http.ResponseWriter,  r *http.Request) {
	
}

// 删除任务
func handlerJobDelete(Res http.ResponseWriter,  Req *http.Request) {
	// var (
	// 	err error
	// )
	// if err = Res. {
		
	// }
}
// 单例对象
var (
	G_apiServe *ApiServer
)

func InitApiServer() (err error){
	var (
		mux *http.ServeMux
		listerner net.Listener
	)
	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handlerJobSave)
	mux.HandleFunc("/job/delete", handlerJobDelete)
	// 监听端口
	listerner, err = net.Listen("tcp", ":"+strconv.Itoa(Config.port))
	if err != nil {
		return err
	}
	// 创建http服务
	httpServer := &http.Serve{
		ReadTimeout: time.Duration(G_apiServe),
		WriteTimeout: time.Duration(G_apiServe),
		Handler:mux,
	}
	// 赋值单例
	G_apiServe  = &ApiServer{
		httpServer: httpServer,
	}
	// 启动服务端
	go httpServer.sever(listerner)
	return nil
}