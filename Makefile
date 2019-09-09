.PHONY: build
build:
	docker-compose -f ./docker/docker-compose.yml build

.PHONY: destroy
destroy:
	docker-compose -f docker/docker-compose.yml down
	docker volume prune

.PHONY: seed
seed:
	./scripts/init-elasticsearch.sh

.PHONY: start
start:
	docker-compose -f ./docker/docker-compose.yml up -d
