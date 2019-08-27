package main

import (
	"net/http"
	"os"

	"github.com/harkce/engine/server"
	"github.com/rs/cors"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/harkce/engine/.env")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})

	handler := c.Handler(server.Router())
	httpServer := &http.Server{Addr: ":8017", Handler: handler}

	println("[server] listening on :8017")
	httpServer.ListenAndServe()
}
