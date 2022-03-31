package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, ctx := errgroup.WithContext(ctx)
	mux := http.NewServeMux()
	server := http.Server{Addr: ":807", Handler: mux}
	group.Go(func() error {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println("server 1")
		})
		err := server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		fmt.Println("http server stop")
		return server.Shutdown(ctx)
	})

	group.Go(func() error {
		c := make(chan os.Signal, 0)
		signal.Notify(c)

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-c:
				fmt.Println("signal", sig)
				cancel()
			}
		}
	})

	err := group.Wait()
	if err != nil {
		fmt.Println(err)
		fmt.Println(ctx.Err())
	}

	fmt.Println("group done")
}
