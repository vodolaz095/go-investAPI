package main

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/vodolaz095/go-investAPI/investapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/*
В этом примере мы с нуля настраиваем GRPC клиент общепринятыми практиками,
потом этим клиентом мы соединяемся с API Тинькофф Инвестициями
*/

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
		log.Fatalf("%s : при соединениии с api", err)
	}
	defer conn.Close()

	instrumentsAPI := investapi.NewInstrumentsServiceClient(conn)

	req := investapi.InstrumentRequest{
		IdType: investapi.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     "BBG00RRT3TX4",
	}

	res, err := instrumentsAPI.GetInstrumentBy(context.TODO(), &req)
	if err != nil {
		log.Fatalf("%s : при поиске инструмента", err)
	}
	log.Printf("Инструмент %s найден!", res.Instrument.Name)
}
