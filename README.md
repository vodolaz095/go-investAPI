Go-InvestAPI
=============================
[![PkgGoDev](https://pkg.go.dev/github.com/vodolaz095/go-investAPI/investapi)](https://pkg.go.dev/github.com/vodolaz095/go-investAPI/investapi?tab=doc)

Неофициальный минималистичный grpc клиент для работы с [investAPI](https://russianinvestments.github.io/investAPI/) платформы
[Т-Банк Инвестиции](https://www.tbank.ru/sl/AugaFvDlqEP)

Предупреждение
===============================
Это не официальная библиотека (хотя, она упомянута в документации к [API T-Банк Инвестиций](https://russianinvestments.github.io/investAPI/faq_golang/)), 
её работоспособность и безопасность не гарантирована автором! Несмотря на то, что большая часть кода библиотеки сгенерирована из официальных 
[определений протокола взаимодействия investAPI в виде proto файлов](https://github.com/RussianInvestments/investAPI/tree/main/src/docs/contracts),
код не проходил независимый аудит безопасности, и не был протестирован представителями банка.
Код может содержать уязвимости и логические ошибки, из-за которых Вы можете потерять деньги.
Используйте этот код на свой страх и риск.


Ссылка на официальный клиент - https://github.com/RussianInvestments/invest-api-go-sdk.
Мой клиент был опубликован раньше, и он субъективно проще.

Пример использования
===============================

1. Устанавливаем зависимости

```shell

$ go get -u github.com/vodolaz095/go-investAPI/investapi 

```

2. Запускаем код
```go

package main

import (
   "context"
   "log"
   "time"

   "google.golang.org/grpc/status"
   
   "github.com/vodolaz095/go-investAPI/investapi"
)

const token = "тутДолженБытьТокен" // https://russianinvestments.github.io/investAPI/grpc/#tinkoff-invest-api_2

func main() {
   client, err := investapi.New(token)
   if err != nil {
      log.Fatalf("%s : при соединении с investAPI", err)
   }
   res, err := client.MarketDataServiceClient.GetLastPrices(context.Background(),
      &investapi.GetLastPricesRequest{Figi: []string{"BBG00RRT3TX4"}},
   )
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

```

3. В директории [example](https://github.com/vodolaz095/go-investAPI/tree/master/example) лежат примеры использования


Как сгенерировать код клиента из proto файлов?
===============================
Официальное описание протокола взаимодействия опубликовано тут [https://github.com/RussianInvestments/investAPI/tree/main/src/docs/contracts](https://github.com/RussianInvestments/investAPI/tree/main/src/docs/contracts).
Код в моём репозитории генерирует клиент для Golang на основе этих `proto` файлов.
Код проверялся на дистрибутиве `Fedora 39+` с `docker` версии 20.10.15, `podman` версии 4.9.4 и `Golang` версии 1.22.8

1. Удалите все файлы из поддиректории `investapi/`, кроме `client.go`, `helpers.go` и `helpers_test.go`!

2. Запустите команду (в зависимости от настройки системы могут потребоваться права суперпользователя для запуска Docker)

```shell

./generate.sh

```
3. Команда должна создать локальный докер образ со всем инструментарием, необходимым для генерации файла клиента
   из *.proto файлов, загруженных отсюда https://github.com/RussianInvestments/investAPI/tree/main/src/docs/contracts

4. При запуске докер образа, контейнер клонирует код репозитория https://github.com/RussianInvestments/investAPI/ и генерирует
   файлы для Go модуля используя `protoc` и прочие инструменты - см `entrypoint.sh` как это происходит.

5. Как результат выполнения предыдущей команды, в поддиректории `investapi/` появятся файлы

```
common.pb.go
instruments.pb.go
instruments_grpc.pb.go
marketdata.pb.go
marketdata_grpc.pb.go
operations.pb.go
operations_grpc.pb.go
orders.pb.go
orders_grpc.pb.go
sandbox.pb.go
sandbox_grpc.pb.go
stoporders.pb.go
stoporders_grpc.pb.go
users.pb.go
users_grpc.pb.go

```

На данный момент файлы включены в код репозитория.


Где используется?
===============

- https://github.com/vodolaz095/stocks_broadcaster - утилита, которая ретранслирует кодировки в redis

Интеграция с OpenTelemetry
================

Основано на примере https://github.com/uptrace/opentelemetry-go-extra/blob/main/example/grpc/client/client.go

```golang
package main

import (
  "google.golang.org/grpc"
  "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
  "github.com/vodolaz095/go-investAPI/investapi"  
)

const token = "тутДолженБытьТокен" // https://tinkoff.github.io/investAPI/grpc/#tinkoff-invest-api_1

func main() {
   // other code
	client, err := investapi.New(token,
      grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
      grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
   )
   // other code
}


```



Лицензия
===============

Все права защищены (c) 2022 Остроумов Анатолий Александрович

Данная лицензия разрешает лицам, получившим копию данного программного обеспечения и сопутствующей документации
(в дальнейшем именуемыми «Программное обеспечение»), безвозмездно использовать Программное обеспечение без ограничений,
включая неограниченное право на использование, копирование, изменение, слияние, публикацию, распространение,
сублицензирование и/или продажу копий Программного обеспечения, а также лицам,
которым предоставляется данное Программное обеспечение, при соблюдении следующих условий:

Указанное выше уведомление об авторском праве и данные условия должны быть
включены во все копии или значимые части данного Программного обеспечения.

ДАННОЕ ПРОГРАММНОЕ ОБЕСПЕЧЕНИЕ ПРЕДОСТАВЛЯЕТСЯ «КАК ЕСТЬ», БЕЗ КАКИХ-ЛИБО ГАРАНТИЙ, ЯВНО ВЫРАЖЕННЫХ ИЛИ ПОДРАЗУМЕВАЕМЫХ,
ВКЛЮЧАЯ ГАРАНТИИ ТОВАРНОЙ ПРИГОДНОСТИ, СООТВЕТСТВИЯ ПО ЕГО КОНКРЕТНОМУ НАЗНАЧЕНИЮ И ОТСУТСТВИЯ НАРУШЕНИЙ,
НО НЕ ОГРАНИЧИВАЯСЬ ИМИ. НИ В КАКОМ СЛУЧАЕ АВТОРЫ ИЛИ ПРАВООБЛАДАТЕЛИ НЕ НЕСУТ ОТВЕТСТВЕННОСТИ ПО КАКИМ-ЛИБО ИСКАМ,
ЗА УЩЕРБ ИЛИ ПО ИНЫМ ТРЕБОВАНИЯМ, В ТОМ ЧИСЛЕ, ПРИ ДЕЙСТВИИ КОНТРАКТА, ДЕЛИКТЕ ИЛИ ИНОЙ СИТУАЦИИ,
ВОЗНИКШИМ ИЗ-ЗА ИСПОЛЬЗОВАНИЯ ПРОГРАММНОГО ОБЕСПЕЧЕНИЯ ИЛИ ИНЫХ ДЕЙСТВИЙ С ПРОГРАММНЫМ ОБЕСПЕЧЕНИЕМ.


Copyright (c) 2021 Остроумов Анатолий Александрович

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
associated documentation files (the "Software"), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge, publish, distribute,
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT
LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
