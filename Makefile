.PHONY: build push run clean

IMAGE_NAME ?= ghcr.io/xusenlin/forest-blog
TAG ?= latest

build:
	docker build -t $(IMAGE_NAME):$(TAG) .

run:
	docker run -d -p 8080:80 --name forest-blog $(IMAGE_NAME):$(TAG)

push:
	docker push $(IMAGE_NAME):$(TAG)

clean:
	docker rm -f forest-blog 2>/dev/null || true
	docker rmi $(IMAGE_NAME):$(TAG) 2>/dev/null || true
