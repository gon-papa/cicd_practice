FROM golang:1.23

# Airが使うバイナリのPATHを通す
ENV PATH="/go/bin:${PATH}"

# 作業ディレクトリ
WORKDIR /app

# Airをインストール
RUN go install github.com/air-verse/air@latest

# tmpディレクトリを先に作成（Airが出力する場所）
RUN mkdir -p tmp

# 必要ファイルをコピー
COPY . .

# Airで起動
CMD ["air", "-c", ".air.toml"]