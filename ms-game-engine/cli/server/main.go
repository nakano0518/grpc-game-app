package main

import (
	"flag"
	grpcSetup "grpc-game-app/ms-game-engine/internal/server/grpc"

	"github.com/rs/zerolog/log"
)

func main() {
	var addessPtr = flag.String("address", ":60051", "address where can connect with ms-gameengine service")
	flag.Parse()

	s := grpcSetup.NewServer(*addessPtr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start grpc server of ms-gameengine")
	}

}
