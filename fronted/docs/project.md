# 项目搭建流程

## 版本

node v16.14.2
pnpm v8.11.0

```shell
# 安装 nvm 管理 node 版本
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash
nvm install 16.14.2
nvm use v16.14.2
# 安装 pnpm
npm i -g pnpm@8.0.0
```

## vscode 插件

vue vscode snippets

## 初始化项目

```shell
pnpm create vite
# project name: fronted
# select vue and TypeScript

cd fronted
pnpm i
pnpm run dev
```
