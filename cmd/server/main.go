package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/BorisGujvin/gin-api/controller"
    "github.com/BorisGujvin/gin-api/content"
)

func main() {

    port := ":80"

    engine := gin.New()

    server := &http.Server{
        Addr:    port,
        Handler: engine,
    }

    engine.GET("/healthcheck", controller.IsHealthy)
    v1 := engine.Group("/api/v1")
    {
        v1.GET("/contents", content.GetContents)
        v1.POST("/contents", content.PostContents)
    }

    chanErrors := make(chan error)
    go func() {
        chanErrors <- runServer(server)
    }()

    chanSignals := make(chan os.Signal, 1)
    signal.Notify(chanSignals, os.Interrupt, syscall.SIGTERM)

    select {
    case err := <-chanErrors:
        log.Fatalf("Error while starting server %s", err)
    case s := <-chanSignals:
        log.Printf("Shutting down server in few seconds due to %s", s)
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        if err := Close(ctx, server); err != nil {
            log.Fatal("Server forced to shutdown: ", err)
        }
        log.Print("Server exiting gracefully")
    }

}

func runServer(server *http.Server) error {
    return server.ListenAndServe()
}
func Close(ctx context.Context, server *http.Server) error {
    return server.Shutdown(ctx)
}