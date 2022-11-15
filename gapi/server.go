package gapi

import (
	"fmt"

	db "github.com/wizlif/simplebank/db/sqlc"
	"github.com/wizlif/simplebank/pb"
	"github.com/wizlif/simplebank/token"
	"github.com/wizlif/simplebank/util"
)

// Server serves gRPC requests for our banking requests
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	db         db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}

	server := &Server{
		config:     config,
		db:         store,
		tokenMaker: tokenMaker,
	}

	
	return server, nil
}
