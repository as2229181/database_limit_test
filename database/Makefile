include .env
export
ifndef CONATINER_PREFIX
$(error CONATINER_PREFIX is not set)
endif

PROJECT_NAME=postgres

IMAGE_NAME=$(CONATINER_PREFIX)/$(PROJECT_NAME):$(VERSION)
CONTAINER_NAME=$(CONATINER_PREFIX)-$(PROJECT_NAME)

BUILD_DIR=_build
BUILD_FILES= \
	${BUILD_DIR}/Dockerfile \
	${BUILD_DIR}/files/docker-entry-point-initdb.d/10-record.sh \
	${BUILD_DIR}/files/docker-entry-point-initdb.d/10-record.sql.template \

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
	docker volume rm pg-data

.PHONY: run
run: build
	docker run -d --name $(CONTAINER_NAME) \
		-v pg-data:/var/lib/postgresql/data \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e RECORD_PASSWORD=$(RECORD_PASSWORD) \
		-p 5432:5432 \
		$(IMAGE_NAME)

$(BUILD_DIR)/Dockerfile: Dockerfile
	mkdir -p $(dir $@)
	cp $< $@

$(BUILD_DIR)/files/%: files/%
	mkdir -p $(dir $@)
	cp $< $@