#!/bin/bash

RUN_ENV=${1:-dev}

if [ $RUN_ENV = "dev" ]; then
    echo "Running in dev mode"
    cp environments/.env.dev .env
else
    echo "Running in prod mode"
    cp environments/.env.prod .env
fi