# api ディレクトリ

api の使用を格納

# OpenAPI の環境構築

## redocly のインストール

```
npm i -g @redocly/cli@latest
```

## 分割された yml を一つに統合。

```
redocly bundle api/root.yaml --output api/api.gen.yaml
```

or

```
npx @redocly/cli bundle api/root.yaml --output api/api.gen.yaml
```

## 公式のコマンド集

https://redocly.com/docs/cli/commands

# 静的解析

```
redocly lint root.yaml
```

# api の仕様を更新した時のコマンド

```
go generate
go build ./...
```
