package main

import (
	"context"
	service2 "demo-golang/homework/shi2zhan4/week01/service"
	"log"
	"net/http"
	"time"
)

func main() {
	s1 := service2.NewKNHTTPServer("9505", "127.0.0.1:9505")
	s1.Handler("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("127.0.0.1:9505 say hello")
	}))
	s2 := service2.NewKNHTTPServer("9506", "127.0.0.1:9506")
	s2.Handler("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("127.0.0.1:9506 say hello")
	}))
	svc := service2.NewKNService([]*service2.KNHTTPServer{s1, s2}, service2.WithShutdownCallback(cacheToDBCallback))
	svc.Start()
}

func cacheToDBCallback(ctx context.Context) {
	chanSgn := make(chan struct{}, 1)
	go func() {
		log.Println("刷新缓存中。。。")
		time.Sleep(1 * time.Second)
		chanSgn <- struct{}{}
	}()
	select {
	case <-chanSgn:
		log.Println("缓存刷新到 DB")
	case <-ctx.Done():
		log.Println("缓存刷新超时")
	}
}