package main

import (
	"context"
	"crypto/tls"
	"github.com/vodolaz095/go-investAPI/investapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
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
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName: "invest-public-api.tinkoff.ru",
		})), // по умолчанию проверяет сертификат
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
	for _, p := range res.LastPrices {
		log.Printf("%s - %s %s\n", p.Figi, p.Price.String(), p.Time.AsTime().Format(time.Stamp))
	}
}
