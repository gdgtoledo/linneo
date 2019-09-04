.PHONY: destroy
destroy:
	docker-compose -f docker/docker-compose.yml down
	docker volume prune

.PHONY: start
start:
	./scripts/init-elasticsearch.sh