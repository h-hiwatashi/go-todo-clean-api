# alpineはバージョンを実質的に固定にできないためDebianベースを使用する
# Builder
FROM golang:1.24-bullseye as builder
# アップデートとgitのインストール
# goを扱うにはgitが必須であるがalpyneやDebianは軽量化のためgitが入っていない
RUN apt update && apt install -y git

WORKDIR /app

# COPY . .

# RUN make build

# Distribution
# FROM alpine:latest

# RUN apt update && apt upgrade && apt install -y git && \
# apt --update --no-cache add tzdata && \
#     mkdir /app

# WORKDIR /app

# EXPOSE 9090

# COPY --from=builder /app/engine /app/

# CMD /app/engine