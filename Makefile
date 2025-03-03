include .env
export $(shell sed 's/=.*//' .env)

PORT ?= $(PORT)

db:
	echo "docker run --rm --name pg-test ..."
	@docker run --rm --name pg-test \
      -e POSTGRES_USER=$(USER) \
      -e POSTGRES_PASSWORD=$(PASSWORD) \
      -e POSTGRES_DB=$(DB) \
      -p 5432:5432 \
      -d postgres:latest

stop_db:
	docker stop pg-test || docker rm pg-test
	docker ps -a


run:
	cd cmd/api  && go run ./... -port=$(PORT)
