compose-build-up:
	make compose-down
	docker-compose -f docker-compose.yaml build;
	docker-compose -f docker-compose.yaml up;

compose-down:
	docker-compose -f docker-compose.yaml down -v;

