package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)



const (
	authTokenKey string = "auth_token"
    authTokenValue string = "authd"
)

func validateAuthToken(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)

	if t, ok := md[authTokenKey]; ok {
		switch {
		case len(t) != 1:
			return status.Errorf(codes.InvalidArgument, "auth_token should conatin only 1 value")
		case t[0] != "authd":
			return status.Errorf(codes.Unauthenticated, "incorrect auth_token")
		}
	} else {
		return status.Errorf(codes.Unauthenticated, "failed to get auth_token")
	}

	return nil
}

func unaryAtuhToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := validateAuthToken(ctx); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func streamAuthToken(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, hanler grpc.StreamHandler) error {
	if err := validateAuthToken(ss.Context()); err != nil {
		return err
	}
	return hanler(srv, ss)
}


func unaryLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println(info.FullMethod, "called")
	return handler(ctx, req)
}

func streamLog(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, hanler grpc.StreamHandler) error {
	log.Println(info.FullMethod, "called")
	return hanler(srv, ss)
}