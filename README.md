# ragingcd

<h1 align="center">
    <img src="./fronted/public/logo.png" alt="Logo">
</h1>
<p align="center">
    <a href="https://github.com/vuejs/vue">
      <img src="https://img.shields.io/badge/vue-3.3.8-brightgreen" alt="vue">
    </a>
    <a href="https://github.com/ElemeFE/element">
      <img src="https://img.shields.io/badge/element--plus-2.4.2-brightgreen" alt="element-plus">
    </a>
    <a href="#">
      <img src="https://img.shields.io/badge/java-1.8-brightgreen" alt="element-plus">
    </a>
    <a href="#">
      <img src="https://img.shields.io/badge/spring--boot-2.7.7-brightgreen" alt="element-plus">
    </a>
    <a href="#">
      <img src="https://img.shields.io/badge/golang-1.20-brightgreen" alt="element-plus">
    </a>
    <a href="#">
        <img src="https://img.shields.io/github/license/jaronnie/ragingcd">
    </a>
</p>

![image-20240104110349987](https://oss.jaronnie.com/ragingcd.png)

## 一键部署

```shell
docker-compose -f docker-compose-production.yaml up -d
```

## 本地开发

```shell
# 安装 task 工具
go install github.com/go-task/task/v3/cmd/task@latest

docker-compose -f docker-compose-dev.yaml up -d
```
