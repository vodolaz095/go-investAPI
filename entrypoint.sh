#!/usr/bin/env bash

# Данный скрипт клонирует репозиторий c официальными proto файлами,
# и на их основе генерирует исходных код клиента для Golang
# в директории /opt/client

set -e

# Убедимся, что компилятор Golang установлен
which go
go version
go env

# Убедимся, что утилита protoc установлена
which protoc
protoc --version

# Клонируем исходный код проекта
cd /opt/
git clone https://github.com/Tinkoff/investAPI.git
ls -l /opt/investAPI/

# Смотрим, крайний коммит, на основе которого мы будет генерировать исходный код клиента
cd /opt/investAPI
git log -1

# Генерируем код модуля
cd /opt/investAPI/src/docs/contracts
protoc \
  --proto_path=/usr/bin/include/google/ \
  --proto_path=/opt/investAPI/src/docs/contracts \
  --go_out=/opt/client \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_out=/opt/client \
  *.proto


# Смотрим, что получилось
ls -l /opt/client/
# Сообщаем, что получилось
echo "Код сгенерирован из коммита https://github.com/Tinkoff/investAPI/commit/$(git log --format='%H' -1)"
