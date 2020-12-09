package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR1)

	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)

	server1 := http.Server{
		Addr: ":8081",
	}

	server2 := http.Server{
		Addr: ":8082",
	}

	http.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
		exit <- syscall.SIGUSR1
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("访问/shutdown 来关闭web服务"))
	})

	group.Go(func() error {
		return server1.ListenAndServe()
	})
	group.Go(func() error {
		return server2.ListenAndServe()
	})

	go func() {
		<-exit
		err1 := server1.Shutdown(ctx)
		if err1 != nil {
			fmt.Println(err1)
		}
		if err2 := server2.Shutdown(ctx); err2 != nil {
			fmt.Println(err2)
		}
	}()

	err := group.Wait()
	if err != nil {
		fmt.Printf("服务结束:%v\n", err)
	}

}
