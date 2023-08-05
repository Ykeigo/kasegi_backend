FROM golang:1.20-bullseye
#FROM golang:1.14-alpine

# コンテナログイン時のディレクトリ指定
WORKDIR /opt/sandbox-docker-compose-go

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY ./src .
# ADD . .

# ビルド
RUN go build -o app main.go
RUN go install github.com/volatiletech/sqlboiler/v4@latest
# 起動
CMD ["/opt/sandbox-docker-compose-go/app"]

EXPOSE 80/tcp