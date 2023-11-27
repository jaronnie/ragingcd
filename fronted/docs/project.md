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

## 配置

### eslint

```shell
pnpm i eslint -D
# 生成 eslint 配置文件
npx eslint --init

pnpm install -D eslint-plugin-import eslint-plugin-vue eslint-plugin-node eslint-plugin-prettier eslint-config-prettier eslint-plugin-node @babel/eslint-parser
```

### stylelint

```shell
pnpm add sass sass-loader stylelint postcss postcss-scss postcss-html stylelint-config-prettier stylelint-config-recess-order stylelint-config-recommended-scss stylelint-config-standard stylelint-config-standard-vue stylelint-scss stylelint-order stylelint-config-standard-scss -D
```

### mock server

```shell
pnpm add vite-plugin-mock@v2.9.6 -D
pnpm add mockjs -D
```
