package common

import (
	"fmt"
	"log/slog"

	"github.com/steve-care-software/gno/gno.land/pkg/log"
	"github.com/steve-care-software/gno/tm2/pkg/commands"
)

// Used to flush the logger.
type logFlusher func()

func LoggerFromServerFlags(serverFlags *ServerFlags, io commands.IO) (*slog.Logger, logFlusher, error) {
	// Initialize the zap logger.
	zapLogger, err := log.InitializeZapLogger(
		io.Out(),
		serverFlags.LogLevel,
		serverFlags.LogFormat,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to initialize zap logger: %w", err)
	}

	// Keep a reference to the zap logger flush function.
	flusher := func() { _ = zapLogger.Sync() }

	// Wrap the zap logger with a slog logger.
	logger := log.ZapLoggerToSlog(zapLogger)

	return logger, flusher, nil
}
