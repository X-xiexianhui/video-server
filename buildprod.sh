#! /bin/bash

#Copy templates
cp -R ./templates ./bin/
mkdir ./bin/videos
# Build web and other services

cd ./work/src/video_server/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd ./work/src/video_server/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd ./work/src/video_server/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd ./work/src/video_server/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web