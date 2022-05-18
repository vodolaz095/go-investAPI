package main

import (
	"context"
	"log"
	"time"

	"github.com/vodolaz095/go-investAPI/investapi"
)

const token = "тутДолженБытьТокен" // https://tinkoff.github.io/investAPI/grpc/#tinkoff-invest-api_1

func main() {
	client, err := investapi.New(token)
	if err != nil {
		log.Fatalf("%s : при соединении с investAPI", err)
	}
	res, err := client.MarketDataServiceClient.GetLastPrices(context.Background(),
		&investapi.GetLastPricesRequest{Figi: []string{"BBG00RRT3TX4"}},
	)
	if err != nil {
		log.Fatalf("%s : при получении котировок инструмента ОФЗ 25084", err)
	}
	for _, price := range res.GetLastPrices() {
		log.Printf("%s назад цена сделки была %.4f рублей",
			time.Now().Sub(price.Time.AsTime()).String(), price.GetPrice().ToFloat64(),
		)
	}
	err = client.Connection.Close()
	if err != nil {
		log.Fatalf("%s : при закрытии соединения", err)
	}
}
