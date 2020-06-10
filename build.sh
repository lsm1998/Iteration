#!/bin/bash

IMAGE_REPOSITORY="ipims"
IMAGE_VERSION="v6"

# 删除镜像
function remove_image() {
  docker images | grep -E "${IMAGE_REPOSITORY}" | awk '{print $3}' | uniq | xargs -I {} docker rmi --force {}
}

# 构建镜像
function build_image() {
  docker build -t $IMAGE_REPOSITORY:${IMAGE_VERSION} .
}

# 停止
function stop() {
  docker rm -f ${IMAGE_REPOSITORY} || true
}

# 启动
function start() {
  docker run -d --name ${IMAGE_REPOSITORY} -p 8088:8088 -v /etc/localtime:/etc/localtime $IMAGE_REPOSITORY:${IMAGE_VERSION}
}

stop

remove_image

build_image

start
