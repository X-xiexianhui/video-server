#!/bin/bash
# Build web and other services
cd $GOPATH/src/video_server/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd $GOPATH/src/video_server/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd $GOPATH/src/video_server/streamServer
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd $GOPATH/src/video_server/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web

#如果videos文件夹不存在，则创建
if [ ! -d ./bin/videos/  ];then
  mkdir ./bin/videos
else
  echo dir exist
fi
#如果template文件不存在，递归拷贝
if [ ! -d ./bin/templates/  ];then
  cp -R ./templates ./bin/
else
  echo dir exist
fi