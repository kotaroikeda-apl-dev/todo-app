# ステージ1: ビルド
FROM golang:1.22-alpine AS builder
WORKDIR /app

# モジュールファイルをコピーして依存関係を取得
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . .
RUN go build -o main .

# ステージ2: 実行環境
FROM alpine:latest

# 必要なパッケージをインストール
RUN apk --no-cache add ca-certificates

# 作業ディレクトリを設定
WORKDIR /root/

# ビルド済みのバイナリをコピー
COPY --from=builder /app/main .

# HTTPS 用の証明書と鍵をコピー
COPY server.crt /root/server.crt
COPY server.key /root/server.key

# ポートを公開
EXPOSE 8080
EXPOSE 8443

# アプリケーションを起動
CMD ["./main"]
