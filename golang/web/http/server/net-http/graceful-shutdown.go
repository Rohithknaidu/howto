// https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("started handler")
	time.Sleep(10*time.Second)
	fmt.Println("completed handler")
	w.Write([]byte("hello world"))
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	// db, err := repo.SetupPostgresDB(ctx, getConfig("DB_DSN", "root@tcp(127.0.0.1:3306)/service"))
	// if err != nil {
	// 	panic(err)
	// }

	mux := http.NewServeMux()
	// assign a route/todo to a handler myHandler
	mux.HandleFunc("/hello", myHandler)

	httpServer := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		fmt.Println("gracefully shutdown........")
		time.Sleep(time.Second)
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}


