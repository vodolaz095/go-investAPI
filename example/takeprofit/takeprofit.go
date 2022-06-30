package main

import (
	"context"
	"log"
	"time"

	"github.com/vodolaz095/go-investAPI/investapi"
)

/*
 * Реализация команды TakeProfit
 * Когда цена крайней сделки по акциям Газпрома достигнет 209.50 рублей, мы ставим лимитную
 * заявку на 209.29 копеек
 * НЕ ЯВЛЯЕТСЯ ИНВЕСТИЦИОННОЙ РЕКОМЕНДАЦИЕЙл
 */

const token = "тутДолженБытьТокен" // https://tinkoff.github.io/investAPI/grpc/#tinkoff-invest-api_1
const accountID = ""               // аккаунт, на который мы покупаем бумагу
const figi = "BBG004730RP0"        // https://www.tinkoff.ru/invest/stocks/GAZP
const minimumPrice = 209.50        // цена, при которой мы ставим лимитную заявку
const wantedPrice = 209.29         // цена лимитной заявки

func main() {
	client, err := investapi.New(token)
	if err != nil {
		log.Fatalf("%s : при соединении с investAPI", err)
	}

	// подписываемся на цену крайней сделки по акциям Газпрома
	request := investapi.MarketDataServerSideStreamRequest{
		SubscribeCandlesRequest:   nil,
		SubscribeOrderBookRequest: nil,
		SubscribeTradesRequest:    nil,
		SubscribeInfoRequest:      nil,
		SubscribeLastPriceRequest: &investapi.SubscribeLastPriceRequest{
			SubscriptionAction: investapi.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
			Instruments: []*investapi.LastPriceInstrument{
				{Figi: figi},
				// тут можно подписаться на цены крайней сделки более чем одной бумаги
			},
		},
	}
	feed := investapi.NewMarketDataStreamServiceClient(client.Connection)
	stream, err := feed.MarketDataServerSideStream(context.TODO(), &request)
	if err != nil {
		log.Fatalf("%s : while subscribbing to feed", err)
	}
	defer stream.CloseSend()

	var msg *investapi.MarketDataResponse
	for {
		msg, err = stream.Recv()
		if err != nil {
			log.Fatalf("%s : while receiving price", err)
		}
		if msg != nil {
			log.Printf("Last prices: %s", msg.String())
			lastPrice := msg.GetLastPrice()

			if lastPrice.GetFigi() == figi {
				if lastPrice.GetPrice() != nil {
					if lastPrice.GetPrice().ToFloat64() < minimumPrice {
						log.Println("Ставим лимитный заказ на покупку бумаги")
						resp, errO := client.OrdersServiceClient.PostOrder(context.TODO(), &investapi.PostOrderRequest{
							Figi:      figi,
							Quantity:  1,
							Price:     investapi.Float64ToQuotation(wantedPrice),
							Direction: investapi.OrderDirection_ORDER_DIRECTION_BUY,
							AccountId: accountID,
							OrderType: investapi.OrderType_ORDER_TYPE_LIMIT,
							OrderId:   "lalala", // тут может быть что угодно для удобства работы с заявками
						})
						if errO != nil {
							log.Printf("%s : while making order", errO)
							break
						}
						log.Printf("Order %s %s is created!", resp.OrderId, resp.Message)
						break // отключаем подписку
					}
				}
			}
			continue
		}
		time.Sleep(10 * time.Millisecond)
	}
}
