#!/bin/bash

docker-compose build --no-cache
docker compose up -d
docker compose logs -t
docker-compose down --rmi all --volumes --remove-orphans

