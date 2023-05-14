package api

import (
	"context"
	"log"
)

type Request struct{}

type Response struct{}

func Main(ctx context.Context, req Request) (Response, error) {
	log.Printf("Hello Morning Night Dream!")

	return Response{}, nil
}
