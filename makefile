# イメージ名
IMAGE_NAME=my-go-dev

# ポート番号
PORT=8080

# Docker イメージをビルド
build:
	docker build -t $(IMAGE_NAME) .

# コンテナを起動（Air + ホットリロード）
run:
	docker run --rm -p $(PORT):8080 -v $(CURDIR):/app $(IMAGE_NAME)

# イメージの削除
clean:
	docker rmi $(IMAGE_NAME) || true

# 再ビルドして起動
rebuild: clean build run