FROM golang:1.20-bullseye
#FROM golang:1.14-alpine

# コンテナログイン時のディレクトリ指定
WORKDIR /opt/sandbox-docker-compose-go

ENV ENV=stg

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY ./src .
# ADD . .

# ビルド
RUN go build -o app main.go auth.go
# 起動
CMD ["/opt/sandbox-docker-compose-go/app"]

EXPOSE 8080/tcp