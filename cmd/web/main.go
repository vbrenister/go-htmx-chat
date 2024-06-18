package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
}

type config struct {
	Addr string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)

	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", app.handleIndex)

	srv := &http.Server{
		Addr:         cfg.Addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err.Error())
		}
	}()

	logger.Info("server started", "port", cfg.Addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)

	logger.Info("shutting down gracefully")
	os.Exit(0)
}
