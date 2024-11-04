package main

import (
	"context"
	"crypto/tls"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/vodolaz095/go-investAPI/investapi"
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
			Token: "тутДолженБытьТокен", // https://russianinvestments.github.io/investAPI/grpc/#tinkoff-invest-api_2
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
		// обработка ошибок
		st, parsed := status.FromError(err)
		if !parsed {
			log.Fatalf("неизвестная ошибка при получении котировок инструмента ОФЗ 25084: %s", err)
		}
		log.Printf("Ошибка получения котировок: %s", st.Err())
		log.Printf("Код ошибки: %s", st.Message())
		log.Printf("Объяснение стандартного кода ошибки: %s", st.Code().String())
		details := st.Details()
		if len(details) > 0 {
			for i := range details {
				log.Printf("Детализация ошибки %v: %q", i, details[i])
			}
		} else {
			log.Printf("Детализации ошибки отсутствует")
		}
		return
	}
	log.Printf("Инструмент %s найден!", res.Instrument.Name)
}
