# note: call scripts from /scripts.PHONY: build up down start stop restart logs ps login

build:
	docker-compose build

up:
	docker-compose up -d

watch:
	docker-compose up -d

down:
	docker-compose down

debug: down
	DEBUG=1 docker-compose up -d

start:
	docker-compose start

stop:
	docker-compose stop

restart: down up

logs:
	docker-compose logs --tail=10 -f

ps:
	docker-compose ps

login: down
	docker-compose run -w /application opina-ai-api /bin/bash