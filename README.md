Go-InvestAPI
=============================
[![PkgGoDev](https://pkg.go.dev/github.com/vodolaz095/go-investAPI/investapi)](https://pkg.go.dev/github.com/vodolaz095/go-investAPI/investapi?tab=doc)

Библиотека для работы с [https://tinkoff.github.io/investAPI/](https://tinkoff.github.io/investAPI/) платформы
[Тинькофф Инвестиции](https://www.tinkoff.ru/sl/AugaFvDlqEP)

Предупреждение
===============================
Это не официальная библиотека, её работоспособность и безопасность не гарантирована автором!
Несмотря на то, что большая часть кода библиотеки сгенерирована из официальных 
[определений протокола взаимодействия investAPI в виде proto файлов](https://github.com/Tinkoff/investAPI/tree/main/src/docs/contracts),
код не проходил независимый аудит безопасности, и не был протестирован представителями банка.
Код может содержать уязвимости и логические ошибки, из-за которых Вы можете потерять деньги.
Используйте этот код на свой страх и риск.


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

```

3. В директории [example](https://github.com/vodolaz095/go-investAPI/tree/master/example) лежат примеры использования


Как сгенерировать код клиента из proto файлов?
===============================
Официальное описание протокола взаимодействия опубликовано тут [https://github.com/Tinkoff/investAPI/tree/main/src/docs/contracts](https://github.com/Tinkoff/investAPI/tree/main/src/docs/contracts).
Код в моём репозитории генерирует клиент для Golang на основе этих `proto` файлов.
Код проверялся на дистрибутиве `Fedora 34` с `docker` версии 20.10.15, `podman` версии 3.4.4 и `Golang` версии 1.16.15

1. Удалите все файлы из поддиректории `investapi/`, кроме `client.go` и `helpers.go`!

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
client.go
helpers.go
helpers_test.go
investapi/common.pb.go
investapi/instruments.pb.go
investapi/instruments_grpc.pb.go
investapi/marketdata.pb.go
investapi/marketdata_grpc.pb.go
investapi/operations.pb.go
investapi/operations_grpc.pb.go
investapi/orders.pb.go
investapi/orders_grpc.pb.go
investapi/sandbox.pb.go
investapi/sandbox_grpc.pb.go
investapi/stoporders.pb.go
investapi/stoporders_grpc.pb.go
investapi/users.pb.go
investapi/users_grpc.pb.go

```

На данный момент файлы включены в код репозитория.


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
