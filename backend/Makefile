include .env
export
ifndef CONATINER_PREFIX
$(error CONATINER_PREFIX is not set)
endif

PROJECT_NAME=backend

IMAGE_NAME=$(CONATINER_PREFIX)/$(PROJECT_NAME):$(VERSION)
CONTAINER_NAME=$(CONATINER_PREFIX)-$(PROJECT_NAME)

BUILD_DIR=_build
BUILD_FILES= \
	${BUILD_DIR}/Dockerfile \
	${BUILD_DIR}/go.mod \
	${BUILD_DIR}/go.sum \
	${BUILD_DIR}/cmd/server/main.go \

.PHONY: all
all: build

.PHONY: build
build: $(BUILD_FILES)
	docker build -t $(IMAGE_NAME) --no-cache\
				--pull\
				${BUILD_DIR}

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	docker rm -f $(CONTAINER_NAME)
	docker rmi $(IMAGE_NAME)

.PHONY: run
run: build
	docker run -d --name $(CONTAINER_NAME) \
		-e HTTP_ADDR=$(HTTP_ADDR) \
		-e RECORD_PASSWORD=$(RECORD_PASSWORD) \
		-p $(HTTP_ADDR):8080 \
		$(IMAGE_NAME)

$(BUILD_DIR)/Dockerfile: Dockerfile
	mkdir -p $(dir $@)
	cp $< $@

$(BUILD_DIR)/go.mod: go.mod
	mkdir -p $(dir $@)
	cp $< $@

$(BUILD_DIR)/go.sum: go.sum
	mkdir -p $(dir $@)
	cp $< $@

$(BUILD_DIR)/cmd/server/main.go: cmd/server/main.go
	mkdir -p $(dir $@)
	cp $< $@
