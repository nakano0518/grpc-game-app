package main

import (
	"context"
	"flag"
	pbhighscore "grpc-game-app/ms-highscore/v1/game"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial ms-highscore gRPC service")
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("adress", *addressPtr).Msg("Failed to close connection")
		}
	}()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c := pbhighscore.NewGameClient(conn)
	if c == nil {
		log.Info().Msg("Client nil")
	}

	r, err := c.GetHighScore(timeoutCtx, &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}
	if r != nil {
		log.Info().Interface("highscore", r.GetHighScore()).Msg("Highscore from ms-highscore microservice")
	} else {
		log.Error().Msg("Couldn't get highscore")
	}

}
