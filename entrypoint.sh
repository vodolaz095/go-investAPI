#!/usr/bin/env bash

# Данный скрипт клонирует репозиторий c официальными proto файлами,
# и на их основе генерирует исходных код клиента для Golang
# в директории /opt/client

set -e

output_dir=/opt

# Убедимся, что компилятор Golang установлен
which go
go version
go env

# Убедимся, что утилита protoc установлена
which protoc
protoc --version

# Клонируем исходный код протофайлов
cd "$output_dir"
git clone https://github.com/RussianInvestments/investAPI.git
ls -l "$output_dir/investAPI/"

# Смотрим, крайний коммит, на основе которого мы будет генерировать исходный код клиента
cd "$output_dir/investAPI"
git log -1

# Сохраняем хэш коммита
githash=$(git log --format='%H' -1)

# Генерируем код модуля
cd "$output_dir/investAPI/src/docs/contracts"
protoc \
  --proto_path=/usr/bin/include/google/ \
  --proto_path="$output_dir/investAPI/src/docs/contracts" \
  --go_out="$output_dir/client" \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_out="$output_dir/client" \
  *.proto


# Генерируем ленивку с хэшем коммита, из которого был сгенерирован клиент
rm -f "$output_dir/client/version.go"
cp "$output_dir/client/version.tpl" "$output_dir/client/version.go"
sed -i "s/development/$githash/g" "$output_dir/client/version.go"

# Смотрим, что получилось
ls -l "$output_dir/client/"

# Сообщаем, что получилось
echo "Код сгенерирован из коммита https://github.com/RussianInvestments/investAPI/commit/$githash"
