include .env
export
ifndef CONATINER_PREFIX
$(error CONATINER_PREFIX is not set)
endif

PROJECT_NAME=backend

IMAGE_NAME=$(CONATINER_PREFIX)/$(PROJECT_NAME):$(VERSION)
CONTAINER_NAME=$(CONATINER_PREFIX)-$(PROJECT_NAME)

BUILD_DIR=_build

.PHONY: all
all: build

.PHONY: stage
stage:
	rm -rf $(BUILD_DIR)
	mkdir -p $(BUILD_DIR)
	cp Dockerfile go.mod go.sum $(BUILD_DIR)/
	cp -r cmd internal migrations $(BUILD_DIR)/

.PHONY: build
build: stage
	docker build -t $(IMAGE_NAME) --no-cache\
				--pull\
				$(BUILD_DIR)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	docker rm -f $(CONTAINER_NAME)
	docker rmi $(IMAGE_NAME)

.PHONY: migrate
migrate:
	go run ./cmd/migrate

.PHONY: run
run: build
	docker run -d --name $(CONTAINER_NAME) \
		-e HTTP_ADDR=$(HTTP_ADDR) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_HOST=$(POSTGRES_HOST) \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-e POSTGRES_PORT=$(POSTGRES_PORT) \
		-p 8080:8080 \
		$(IMAGE_NAME)
