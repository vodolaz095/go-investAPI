Go-InvestAPI
=============================

Библиотека для работы с [https://tinkoff.github.io/investAPI/](https://tinkoff.github.io/investAPI/).

Находится в разработке, так как официальной документации и примеров для Go пока ещё нет
https://github.com/Tinkoff/investAPI/issues/129

Предупреждение
===============================
Это не официальная библиотека, её работоспособность и безопасность не гарантирована автором!
Используйте этот код на свой страх и риск.


Как сгенерировать код клиента из proto файлов?
===============================
Код проверялся на дистрибутиве Fedora 34 с podman версии 3.4.4 и Golang версии 1.16.15

1. Удалите все файлы из под директории `investapi/`

2. Запустите команду 

```shell

./generate.sh

```
3. Команда должна создать локальный докер образ со всем инструментарием, необходимым для генерации файла клиента
   из *.proto файлов, загруженных отсюда https://github.com/Tinkoff/investAPI/tree/main/src/docs/contracts

4. При запуске докер образа, контейнер клонирует код репозитория https://github.com/Tinkoff/investAPI/ и генерирует
   файлы для Go модуля используя `protoc` и прочие инструменты - см `entrypoint.sh` как это происходит.

5. Как результат выполнения предыдущей команды, в поддиректории `investapi/` появятся файлы
```

investapi/common.pb.go
investapi/instruments.pb.go
investapi/marketdata.pb.go
investapi/operations.pb.go
investapi/orders.pb.go
investapi/sandbox.pb.go
investapi/stoporders.pb.go
investapi/users.pb.go

```

На данный момент файлы включены в код репозитория


Пример использования
===============================
Внимание - я не уверен, что сгенерированный код модуля надо использовать именно 
так, но, на данный момент (8 апреля 2022 года) никакой официальной документации и примеров для Go не существует!

```go

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


```
