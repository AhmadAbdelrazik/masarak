package httpserver

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/rs/zerolog/log"
)

func Serve(routes http.Handler, cfg *config.Config) error {
	if cfg == nil {
		panic("missing config")
	}
	srv := http.Server{
		Addr:         cfg.HostURL,
		Handler:      routes,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	shutdownerr := make(chan error, 1)
	go func() {

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		sig := <-quit
		log.Info().Str("signal", sig.String()).Msg("Shutting down the server")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownerr <- err
		}

		shutdownerr <- nil
	}()

	log.Info().Str("Addr", srv.Addr).Str("env", cfg.Enviroment).Msg("Shutting down the server")

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	if err := <-shutdownerr; err != nil {
		return err
	}

	log.Info().Msg("stopped the server")

	return nil
}
