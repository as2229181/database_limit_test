.PHONY: up down logs ps clean
up:
	docker compose up -d --build

down:
	docker compose down

logs:
	docker compose logs -f

ps:
	docker compose ps

clean:
	docker compose down -v

migrate:
	docker compose --profile migrate run --rm migrate

.PHONY: database-build
database-build:
	make -C database

.PHONY: database-run
database-run:
	make -C database run

.PHONY: database-clean
database-clean:
	make -C database clean

.PHONY: database-volume-clean
database-volume-clean:
	docker volume rm pg-data

.PHONY: backend-build
backend-build:
	make -C backend

.PHONY: backend-run
backend-run:
	make -C backend run

.PHONY: backend-clean
backend-clean:
	make -C backend clean
