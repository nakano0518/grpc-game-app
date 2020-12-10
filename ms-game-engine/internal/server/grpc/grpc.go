package grpc

import (
	"context"
	"grpc-game-app/ms-game-engine/internal/server/logic"
	pbgameengine "grpc-game-app/ms-game-engine/v1/gameengine"
	"net"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) GetSize(ctx context.Context, input *pbgameengine.GetSizeRequest) (*pbgameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in ms-game-engine called")
	return &pbgameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

func (g *Grpc) SetScore(ctx context.Context, input *pbgameengine.SetScoreRequest) (*pbgameengine.SetScoreResponse, error) {
	log.Info().Msg("SetScore in ms-game-engine called")
	set := logic.SetScore(input.Score)
	return &pbgameengine.SetScoreResponse{
		Set: set,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "fail to open tcp port")
	}

	serverOpts := []grpc.ServerOption{}
	g.srv = grpc.NewServer(serverOpts...)

	pbgameengine.RegisterGameEngineServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for ms-gameengine microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for ms-gameengine microservice")
	}
	return nil
}
