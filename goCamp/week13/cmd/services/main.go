package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
)

func main() {
	//init background context
	ctx := context.Background()
	//with cancel
	ctx, cancel := context.WithCancel(ctx)
	//with errgroup
	group, ctx := errgroup.WithContext(ctx)
	server := InitializeHttpServer()

	//start http server
	group.Go(func() error {
		err := server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	})

	//handle server shutdown
	group.Go(func() error {
		<-ctx.Done()
		fmt.Println("http server stop")
		return server.Shutdown(ctx)
	})

	group.Go(func() error {
		//register os signal
		c := make(chan os.Signal, 0)
		signal.Notify(c)

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-c:
				fmt.Println("signal", sig)
				//cancel ctx when signal receive
				cancel()
			}
		}
	})

	//opentracing.SetGlobalTracer()
	tracer := opentracing.GlobalTracer()
	parentSpan := tracer.StartSpan("serviceA")
	defer parentSpan.Finish()

	err := group.Wait()
	if err != nil {
		fmt.Println(err)
		fmt.Println(ctx.Err())
	}

	fmt.Println("group done")
}
