package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init gin
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.GET("/hw", helloHandler())

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("can't run webserver")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
