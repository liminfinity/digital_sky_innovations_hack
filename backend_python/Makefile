DOCKER_COMPOSE := docker compose
IMAGE_NAME := backend_python-app:latest

.PHONY: run serve gen tunnel up up-with-logs rm bash build logs

## Запуск FastAPI (локально, с hot-reload)
run:
	poetry run uvicorn app.main:app --reload

## Запуск FastAPI без hot-reload (для билда)
serve:
	poetry run uvicorn app.main:app --host 0.0.0.0 --port 8000

## Генерация OpenAPI схемы
gen:
	poetry run python -m app.gen_openapi

## SSH-туннель
tunnel:
	ssh -p 443 -R0:localhost:8000 qr@a.pinggy.io

## Запуск контейнеров в фоне
up:
	$(DOCKER_COMPOSE) up --detach --force-recreate

## Запуск контейнеров с логами
up-with-logs:
	$(DOCKER_COMPOSE) up --force-recreate

## Остановка и удаление контейнеров + образа
rm:
	$(DOCKER_COMPOSE) stop
	$(DOCKER_COMPOSE) down --remove-orphans
	docker image rm $(IMAGE_NAME) || true

## Подключение внутрь контейнера
bash:
	$(DOCKER_COMPOSE) exec app /bin/bash

## Пересобрать образ
build:
	$(DOCKER_COMPOSE) build --no-cache

## Показать логи
logs:
	$(DOCKER_COMPOSE) logs -f
