package tool

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 优雅退出HTTP服务, server 创建的HTTP服务，timeout 超时时间，sig 需要监听的信号集合，可不填写
func StartAndQuitHTTPGraceful(server *http.Server, timeout time.Duration, sig ...os.Signal) {

	// 启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Fprintln(ErrorHandle, "启动服务失败：", err)
			os.Exit(1)
		}
	}()

	// 创建信号channel
	quit := make(chan os.Signal)
	// 监听信号，若不赋值，则监听所有信号
	signal.Notify(quit, sig...)

	// 阻塞直到输出停止服务信号
	fmt.Fprintln(InfoHandle, "开始停止服务：", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintln(ErrorHandle, "停止服务出错: ", err)
		os.Exit(1)
	}

	fmt.Fprintln(InfoHandle, "服务已经停止")
}
