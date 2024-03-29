package bff

import (
	"context"
	pbgameengine "grpc-game-app/ms-game-engine/v1/gameengine"
	pbhighscore "grpc-game-app/ms-highscore/v1/game"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (gr *gameResource) SetHighScore(c *gin.Context) {
	highScoreString := c.Param("hs")
	highScoreFloat64, err := strconv.ParseFloat(highScoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect highscore to float")
	}
	gr.gameClient.SetHighScore(context.Background(), &pbhighscore.SetHighScoreRequest{
		HighScore: highScoreFloat64,
	})
}

func (gr *gameResource) GetHighScore(c *gin.Context) {
	highScoreResponse, err := gr.gameClient.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Error().Err(err).Msg("Error while getting highscore")
		return
	}
	hsString := strconv.FormatFloat(highScoreResponse.HighScore, 'e', -1, 64)
	c.JSONP(200, gin.H{
		"hs": hsString,
	})
}

func (gr *gameResource) GetSize(c *gin.Context) {
	sizeResponse, err := gr.gameEngineClient.GetSize(context.Background(), &pbgameengine.GetSizeRequest{})
	if err != nil {
		log.Error().Err(err).Msg("Error while getting size")
	}
	c.JSON(200, gin.H{
		"size": sizeResponse.GetSize(),
	})
}

func (gr *gameResource) SetScore(c *gin.Context) {
	scoreString := c.Param("score")
	scoreFloat64, _ := strconv.ParseFloat(scoreString, 64)

	_, err := gr.gameEngineClient.SetScore(context.Background(), &pbgameengine.SetScoreRequest{
		Score: scoreFloat64,
	})
	if err != nil {
		log.Error().Err(err).Msg("Error while setting score in m-game-engine")
	}
}
