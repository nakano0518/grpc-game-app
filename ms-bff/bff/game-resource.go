package bff

import (
	pbgameengine "grpc-game-app/ms-game-engine/v1/gameengine"
	pbhighscore "grpc-game-app/ms-highscore/v1/game"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type gameResource struct {
	gameClient       pbhighscore.GameClient
	gameEngineClient pbgameengine.GameEngineClient
}

func NewGameResource(gameClient pbhighscore.GameClient, gameEngineClient pbgameengine.GameEngineClient) *gameResource {
	return &gameResource{
		gameClient:       gameClient,
		gameEngineClient: gameEngineClient,
	}
}

func NewGrpcGameServiceClient(serverAddr string) (pbhighscore.GameClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
		return nil, err
	} else {
		log.Info().Msgf("Successfullyconnected to [%s]", serverAddr)
	}
	if conn == nil {
		log.Info().Msg("ms-highscore connection is nil in ms-bff")
	}
	client := pbhighscore.NewGameClient(conn)
	return client, nil
}

func NewGrpcGameEngineServiceClient(serverAddr string) (pbgameengine.GameEngineClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
		return nil, err
	} else {
		log.Info().Msgf("Successfullyconnected to [%s]", serverAddr)
	}
	if conn == nil {
		log.Info().Msg("ms-game-engine connection is nil in ms-bff")
	}
	client := pbgameengine.NewGameEngineClient(conn)
	return client, nil
}
