package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authTokenKey   string = "auth_token"
	authTokenValue string = "authd"
	grpcService           = 5 // grpc.service
	grpcMethod            = 7 // grpc.method
)

func validateAuthToken(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.DataLoss, "Not arguments found")
	}

	if t, ok := md[authTokenKey]; ok {
		switch {
		case len(t) != 1:
			return nil, status.Errorf(codes.InvalidArgument, "auth_token should conatin only 1 value")
		case t[0] != "authd":
			return nil, status.Errorf(codes.Unauthenticated, "incorrect auth_token")
		}
	} else {
		return nil, status.Errorf(codes.Unauthenticated, "failed to get auth_token")
	}

	return ctx, nil
}

func unaryLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println(info.FullMethod, "called")
	return handler(ctx, req)
}

func streamLog(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, hanler grpc.StreamHandler) error {
	log.Println(info.FullMethod, "called")
	return hanler(srv, ss)
}

func logCalls(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, level logging.Level, msg string, fields ...any) {
		switch level {
		case logging.LevelDebug:
			msg = fmt.Sprintf("DEBUG :%v", msg)
		case logging.LevelInfo:
			msg = fmt.Sprintf("INFO :%v", msg)
		case logging.LevelWarn:
			msg = fmt.Sprintf("WARM :%v", msg)
		case logging.LevelError:
			msg = fmt.Sprintf("ERROR :%v", msg)
		default:
			panic(fmt.Sprintf("unknown level %v", level))
		}

		l.Println(msg, fields[grpcService], fields[grpcMethod])
	})
}
