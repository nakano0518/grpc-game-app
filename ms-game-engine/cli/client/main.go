package main

import (
	"context"
	"flag"
	pbgameengine "grpc-game-app/ms-game-engine/v1/gameengine"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:60051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial ms-game-engine gRPC service")
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("adress", *addressPtr).Msg("Failed to close connection")
		}
	}()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c := pbgameengine.NewGameEngineClient(conn)
	if c == nil {
		log.Info().Msg("Client nil")
	}

	r, err := c.GetSize(timeoutCtx, &pbgameengine.GetSizeRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}
	if r != nil {
		log.Info().Interface("size", r.GetSize()).Msg("Highscore from ms-game-engine microservice")
	} else {
		log.Error().Msg("Couldn't get size")
	}

}
