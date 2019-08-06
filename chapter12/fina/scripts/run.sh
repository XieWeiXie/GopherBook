#!/usr/bin/env bash

docker-compose up -d -f ../deploements/Docker-compose

mysql -uroot -p < ../data/fina-2019-08-06.sql

