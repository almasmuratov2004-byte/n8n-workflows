package logger

import (
	"os"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	_ = os.MkdirAll("logs", os.ModePerm)

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot open log file")
	}
	multi := zerolog.MultiLevelWriter(os.Stdout, file)

	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	log.Info().Msg("logger initialized")
}
