#!/bin/bash

RUN_ENV=${1:-dev}

if [ $RUN_ENV = "dev" ]; then
    echo "Running in dev mode"
    cp environments/.env.dev .env
    cp environments/.env.dev ./cmd/api/application/.env
    cp environments/.env.dev ./cmd/api/services/.env
else
    echo "Running in prod mode"
    cp environments/.env.prod .env
    cp environments/.env.prod ./cmd/api/application/.env
    cp environments/.env.prod ./cmd/api/services/.env
fi