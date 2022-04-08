#!/usr/bin/env bash

set -e

which go
go version

go env

which protoc
protoc --version

# Клонируем исходный код проекта
cd /opt/
git clone https://github.com/Tinkoff/investAPI.git
ls -l /opt/investAPI/

# Генерируем код модуля
cd /opt/investAPI/src/docs/contracts
/usr/bin/protoc --proto_path=/usr/bin/include/google/ --proto_path=/opt/investAPI/src/docs/contracts --go_out=/opt/client *.proto

# Смотрим, что получилось
ls -l /opt/client/
