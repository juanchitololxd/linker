#!/bin/bash

PID=$(lsof -t -i :8080)

if [-n "$PID"]; then
    kill $PID
    echo "Proceso con PID $PID detenido."
else
    echo "No se encontro ningun proceso en el puerto 8080."
fi