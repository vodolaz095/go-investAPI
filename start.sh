#!/usr/bin/env bash

docker build -t localhost/go-invest-api .

docker run -v `pwd`/client:/opt/client/ localhost/go-invest-api
