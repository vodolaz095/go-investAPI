package main

import (
	"context"
	"log"

	"github.com/vodolaz095/go-investAPI/investapi"
)

/*
Данный пример создаёт тестовый аккаунт и пополняет его 100 рублями
*/

const token = "тутДолженБытьТокен" // https://russianinvestments.github.io/investAPI/grpc/#tinkoff-invest-api_2

func main() {
	client, err := investapi.NewWithCustomEndpoint(token, investapi.SandboxEndpoint)
	if err != nil {
		log.Fatalf("%s : при соединении с investAPI", err)
	}
	sandboxService := investapi.NewSandboxServiceClient(client.Connection)

	resp, err := sandboxService.OpenSandboxAccount(context.Background(), &investapi.OpenSandboxAccountRequest{})
	if err != nil {
		log.Fatalf("%s : while opening sandbox account", err)
	}
	log.Printf("Тестовый счёт %s открыт", resp.AccountId)

	sandboxAccounts, err := sandboxService.GetSandboxAccounts(context.Background(), &investapi.GetAccountsRequest{})
	if err != nil {
		log.Fatalf("%s : while getting sandbox accounts", err)
	}
	var howMuch investapi.MoneyValue
	howMuch = *investapi.MoneyValueFromFloat64(100)
	howMuch.Currency = "RUB"
	for _, account := range sandboxAccounts.GetAccounts() {
		if resp.AccountId == account.Id {
			_, err = sandboxService.SandboxPayIn(context.Background(), &investapi.SandboxPayInRequest{
				AccountId: account.Id,
				Amount:    &howMuch,
			})
			if err != nil {
				log.Fatalf("%s : при попытке пополнить тестовый счёт %s", err, account.Id)
			}
			log.Printf("Тестовый аккаунт %s пополнен на 100 рублей", account.Id)
		}
	}
}
