package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ahmedelg3brii/chat-app/foundation/logger"
	//"github.com/ahmedelg3brii/chat-app/tree/main/foundation/otel"
)

func main() {

	var log *logger.Logger

	traceIDFn := func(ctx context.Context) string {
		return "" // otel.GetTraceID(ctx)
	}

	log = logger.New(os.Stdout, logger.LevelInfo, "cap", traceIDFn)

	// -------------------------------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	log.Info(ctx, "startup", "status", "starting service")
	defer log.Info(ctx, "shutdown", "status", "shutdown complete")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	// -------------------------------------------------------------------------
	// Start/Stop

	return nil
}
