#!/bin/bash

cd ~/linker-prod
nohup ./url-shortener > ./url-shortener.log 2>&1 &
echo "App running in background."