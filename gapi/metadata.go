package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type MetaData struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetaData(ctx context.Context) *MetaData {
	mtdt := &MetaData{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// log.Printf("md %v\n", md)

		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if clientIps := md.Get(xForwardedForHeader); len(clientIps) > 0 {
			mtdt.ClientIP = clientIps[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}
