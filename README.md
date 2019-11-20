# fuhsicloud

伏羲平台

fuhsicloud是一个基于了kubernetes的应用管理系统，通过可视化的页面对应用进行管理，降低容器化成本，同时也降低了Docker及Kubernetes的学习门槛。


## 架构设计

该平台提供了一整套解决方案。


## 安装说明

平台后端基于[go-kit](https://github.com/go-kit/kit)、前端基于[ant-design](https://github.com/ant-design/ant-design)(版本略老)框架进行开发。

后端所使用到的依赖全部都在[go.mod](go.mod)里，前端的依赖在`package.json`，详情的请看`yarn.lock`，感谢开源社区的贡献。

### 安装教程



### 依赖

- Golang 1.12+ [安装手册](https://golang.org/dl/)
- MySQL 5.7+ (大多数据都存在mysql)
- Docker 18.x+ [安装](https://docs.docker.com/install/)
- RabbitMQ (主要用于消息队列)
- Jenkins 2.176.2+ (老版本对java适配可能会有问题，尽量使用新版本)

## 快速开始

1. 克隆

```
$ mkdir -p $GOPATH/src/github.com/fuhsicloud
$ cd $GOPATH/src/github.com/fuhsicloud
$ git clone https://github.com/fuhsicloud/fuhsicloud.git
$ cd fuhsicloud
```

2. 配置文件准备

    - 将连接Kubernets的kubeconfig文件放到该项目目录

3. docker-compose 启动

```
$ cd install/docker-compose
$ docker-compose up
```

4. make 启动

```
$ make run
```
