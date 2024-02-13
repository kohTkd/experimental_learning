.PHONY: docker.dev
docker.dev:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml ${cmd} ${opts} ${container} ${args}

.PHONY: docker.api-main-dev docker.api-main-dev.up docker.api-main-dev.stop \
				docker.api-main-dev.restart docker.api-main-dev.reboot
docker.api-main-dev:
	make docker.dev container='api-main-dev'
docker.api-main-dev.up:
	make docker.api-main-dev cmd='up' opts='-d'
docker.api-main-dev.stop:
	make docker.api-main-dev cmd='stop'
docker.api-main-dev.restart:
	make docker.api-main-dev cmd='restart'
docker.api-main-dev.reboot:
	make docker.api-main-dev.stop && make docker.api-main-dev.up

.PHONY: docker.proto docker.proto.build docker.proto.run
docker.proto:
	make docker.dev container='protobuf'
docker.proto.build:
	make docker.proto cmd='build'
docker.proto.run:
	make docker.proto cmd='run' opts='--rm -it'
