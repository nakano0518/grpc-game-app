package main

import (
	"flag"
	grpcSetup "grpc-game-app/ms-highscore/internal/server/grpc"

	"github.com/rs/zerolog/log"
)

func main() {
	var addessPtr = flag.String("address", ":50051", "address where can connect with ms-highscore service")
	flag.Parse() // The key "adress" is accessable

	s := grpcSetup.NewServer(*addessPtr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start grpc server of ms-highscore")
	}

}
