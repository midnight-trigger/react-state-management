migrate.up.%:
	docker compose --env-file .env -f build/docker-compose.yml build core-migrate
	docker compose --env-file .env -f build/docker-compose.yml run --rm core-migrate up $*
migrate.down.%:
	docker compose --env-file .env -f build/docker-compose.yml build core-migrate
	docker compose --env-file .env -f build/docker-compose.yml run --rm core-migrate down $*
migrate.up:
	docker compose --env-file .env -f build/docker-compose.yml build core-migrate
	docker compose --env-file .env -f build/docker-compose.yml run --rm core-migrate up
migrate.cleanup:
	docker compose --env-file .env -f build/docker-compose.yml build core-migrate
	docker compose --env-file .env -f build/docker-compose.yml run --rm core-migrate cleanup
migrate.down:
	docker compose --env-file .env -f build/docker-compose.yml build core-migrate
	docker compose --env-file .env -f build/docker-compose.yml run --rm core-migrate down

run:
	docker compose --env-file .env -f build/docker-compose.yml up
down:
	docker compose --env-file .env -f build/docker-compose.yml down
build: down
	docker compose --env-file .env -f build/docker-compose.yml build
