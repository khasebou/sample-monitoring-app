#!/bin/sh

rootDir=$(pwd)/..

docker build --tag localhost:5001/internal-microservice:latest $rootDir/internal-microservice
docker push localhost:5001/internal-microservice:latest

docker build --tag localhost:5001/public-microservice:latest $rootDir/public-microservice
docker push localhost:5001/public-microservice:latest