package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"

	"github.com/vodolaz095/go-investAPI/investapi"
)

/*
В этом примере мы создаём GRPC клиент с помощью встроенного конструктора,
потом этим клиентом мы соединяемся с API Тинькофф Инвестициями
и получаем цену крайней сделки по облигации федерального займа 25084
https://www.tinkoff.ru/invest/bonds/SU25084RMFS3.
Мы отображаем все хедеры, которые передал сервер.
*/

const token = "тутДолженБытьТокен" // https://tinkoff.github.io/investAPI/grpc/#tinkoff-invest-api_1

func main() {
	client, err := investapi.New(token)
	if err != nil {
		log.Fatalf("%s : при соединении с investAPI", err)
	}
	// в эту переменную мы записываем все хедеры, которые послал для нас сервер брокера
	var header metadata.MD

	// так мы вызываем метод для получения крайней цены сделки
	res, err := client.MarketDataServiceClient.GetLastPrices(context.Background(),
		// так мы посылаем тело запроса
		&investapi.GetLastPricesRequest{Figi: []string{"BBG00RRT3TX4"}},
		// так мы извлекаем все хедеры,
		// которые сервер посылает вместе с ответом на запрос GetLastPrices
		grpc.Header(&header),
	)

	// так мы вызываем метод для получения крайней цены сделки, если хедеры нам не нужны
	//res, err := client.MarketDataServiceClient.GetLastPrices(context.Background(),
	//		&investapi.GetLastPricesRequest{Figi: []string{"BBG00RRT3TX4"}},
	//)

	for k, v := range header {
		log.Printf("Получен хедер %s : %s", k, v)
		//		2022/06/07 10:58:40 Получен хедер date : [Tue, 07 Jun 2022 07:58:39 GMT]
		//		2022/06/07 10:58:40 Получен хедер x-ratelimit-reset : [20]
		//		2022/06/07 10:58:40 Получен хедер x-ratelimit-limit : [300, 300;w=60]
		//		2022/06/07 10:58:40 Получен хедер grpc-accept-encoding : [gzip]
		//		2022/06/07 10:58:40 Получен хедер server : [envoy]
		//		2022/06/07 10:58:40 Получен хедер x-envoy-upstream-service-time : [1]
		//		2022/06/07 10:58:40 Получен хедер content-type : [application/grpc]
		//		2022/06/07 10:58:40 Получен хедер x-tracking-id : [8f261b2d7286743623221465e853c9c5]
		//		2022/06/07 10:58:40 Получен хедер x-ratelimit-remaining : [181]
	}
	if err != nil {
		log.Fatalf("%s : при получении котировок инструмента ОФЗ 25084", err)
	}
	for _, price := range res.GetLastPrices() {
		log.Printf("%s назад цена сделки была %.4f рублей",
			time.Now().Sub(price.Time.AsTime()).String(), price.GetPrice().ToFloat64(),
		)
	}
	//		2022/06/07 10:58:40 1m52.663352891s назад цена сделки была 95.5830 рублей
	err = client.Connection.Close()
	if err != nil {
		log.Fatalf("%s : при закрытии соединения", err)
	}
}
