#!/bin/bash
cd ../../
sudo docker build -f ./build/backend.Dockerfile -t nichyjt/findnus_backendlocal:0.1 .
sudo docker run --network host nichyjt/findnus_backendlocal:0.1