package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	HttpServer *http.Server
	Logger     *log.Logger
}

func NewRouter(logger *log.Logger) *Server {

	router := http.NewServeMux()

	router.HandleFunc("/", handlers.HtmlHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		HttpServer: httpServer,
		Logger:     logger,
	}
}
