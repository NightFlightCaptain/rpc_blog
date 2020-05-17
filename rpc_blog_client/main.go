package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rpc_blog_client/conf"
	"rpc_blog_client/routers"
	"rpc_blog_client/rpc"
	"time"
)

func initGin()  {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Config.Server.HTTPPort),
		Handler:        router,
		ReadTimeout:    time.Duration(conf.Config.Server.ReadTimeout),
		WriteTimeout:   time.Duration(conf.Config.Server.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		//ListenAndServer总是返回一个错误
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//开始关闭
	log.Printf("Shutdown Server... ")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//ctx会等待5秒让Shutdown执行，如果这5秒Shutdown执行完了，那么err就是server自己的错误。
	//如果5秒内没执行完，那么就会返回ctx的错误
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

func main() {

	conf.SetUp()
	rpc.SetUp()
	initGin()
}
