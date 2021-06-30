﻿# DevLab
## Overview
The basic idea is to give a web portal for develper  which can support create VM, k8s, some middle softwares in an very fast and simple way

## Architecture Diagram
![architecture diagram](./views/image/cloudLab.png)

## Features
- Virtual Machine Management(libvirt, kvm)
- Multi-Node Inter-connection(hostgw)
- External Access(iptables dnat)
- Auto Lifecycle Management
- In-Memory Persistant
- Webex Events Notification
- K8s Cluster Management(Todo)
- SaaS Management(Todo)

## Installation
- controller 
#### Build docker image
```
docker build -t controller .
```
#### Run container
```
docker run -d --net host --env HTTPS_PROXY=xxxxx --env NO_PROXY="xxxx" --env BOT_TOKEN=xxxxx -v "$(pwd)"/.db/:/app/.db -v "$(pwd)"/config.ini:/app/config.ini controller
```
- deployer


How to install deployer? refer to [Deployer repo](https://github.com/JinlongWukong/CloudLab-ansible)

## How to use
- Edit config.ini 
- Run docker container
- Open web brower -> http://cloudLab ip:8088/
