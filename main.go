package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	db "github.com/saltchang/magfile-server/db/sqlc"
	"github.com/saltchang/magfile-server/handler"
	"github.com/saltchang/magfile-server/router"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		USE_SSL    = os.Getenv("USE_SSL")
		HOST       = os.Getenv("HOST")
		PORT       = os.Getenv("PORT")
		dbDriver   = os.Getenv("POSTGRES_DRIVER")
		dbUser     = os.Getenv("POSTGRES_USER")
		dbPassword = os.Getenv("POSTGRES_PASSWORD")
		dbHost     = os.Getenv("POSTGRES_HOST")
		dbPort     = os.Getenv("POSTGRES_PORT")
		dbName     = os.Getenv("POSTGRES_DB")
		dbParams   = os.Getenv("POSTGRES_PARAMS")
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	log.Printf("Connecting database: %s", dbDriver)

	database, err := db.Init(dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName, dbParams)

	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	server := &http.Server{}
	h := handler.NewHandler(database)
	server.Handler = router.UseRouter(h)

	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)

	var protocol string

	if USE_SSL == "true" {
		protocol = "https"
	} else {
		protocol = "http"
	}
	log.Printf("Server started on %s://%s:%s", protocol, HOST, PORT)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

// Stop stops the server and return error when server cannot be shut down correctly
func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
