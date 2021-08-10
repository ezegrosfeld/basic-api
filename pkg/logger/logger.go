package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	production bool = os.Getenv("ENVIRONMENT") == "production"
)

func Initialize() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	lvl := zerolog.InfoLevel

	if !production {
		lvl = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(lvl)
}

func New(domain string, ld string) zerolog.Logger {
	return zerolog.New(os.Stderr).With().Str("domain", domain).Str("ld", ld).Timestamp().Logger()
}
