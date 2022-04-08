package main

import (
	"context"
	"github.com/vodolaz095/go-investAPI/investapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type tokenAuth struct {
	Token string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + t.Token,
	}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return true
}

func main() {
	conn, err := grpc.Dial("invest-public-api.tinkoff.ru:443",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(tokenAuth{
			Token: "тутДолженБытьТокен", // https://tinkoff.github.io/investAPI/grpc/#tinkoff-invest-api_1
		}),
	)
	if err != nil {
		log.Fatalf("%s : while connecting to api", err)
	}
	defer conn.Close()

	req := investapi.GetLastPricesRequest{Figi: []string{"SU25084RMFS3"}}
	res := investapi.GetLastPricesResponse{}

	err = conn.Invoke(context.TODO(), "GetLastPrices", &req, &res)
	if err != nil {
		log.Fatalf("%s : while getting last prices", err)
	}
	log.Println("Response", res)
}
