.PHONY: build push run clean

IMAGE_NAME ?= ghcr.io/xusenlin/forest-blog
TAG ?= latest
PLATFORM ?= linux/amd64

build:
	docker buildx build --platform $(PLATFORM) -t $(IMAGE_NAME):$(TAG) --load .

run:
	docker run -d -p 8080:80 --name forest-blog $(IMAGE_NAME):$(TAG)

push: build
	docker push $(IMAGE_NAME):$(TAG)

clean:
	docker rm -f forest-blog 2>/dev/null || true
	docker buildx rm $(IMAGE_NAME) 2>/dev/null || true
	docker rmi $(IMAGE_NAME):$(TAG) 2>/dev/null || true
