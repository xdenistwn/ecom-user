create-network:
	docker network create internal_vpc

up:
	docker compose -f docker-compose.user.yml up -d

stop:
	docker compose -f docker-compose.user.yml stop

down:
	docker compose -f docker-compose.user.yml down