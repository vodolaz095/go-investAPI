#!/usr/bin/env bash

# можно использовать и docker
podman build -t localhost/go-invest-api .
podman run -v `pwd`/investapi:/opt/client/ localhost/go-invest-api

# делаем go.sum файл
go mod tidy
