#!/usr/bin/env bash

set -e

# можно использовать и docker, и podman

containerEngine=docker
# containerEngine=podman


# создаём базовый образ
"$containerEngine" build -t localhost/go-invest-api .

# запускаем скрипт entrypoint.sh в базовом образе
"$containerEngine" run -v `pwd`/investapi:/opt/client/ localhost/go-invest-api

# делаем go.sum файл
go mod tidy
