#!/bin/bash

docker build . -t maozzadok/hello-world
docker push maozzadok/hello-world:latest

