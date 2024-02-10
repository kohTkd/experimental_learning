.PHONY: docker.api-main.build docker.api-main.run docker.api-main.up docker.api-main.stop docker.api-main.restart docker.api-main.reboot
docker.api-main.build:
	make docker.build container='api-main'
docker.api-main.run:
	make docker.run container='api-main'
docker.api-main.up:
	make docker.up container='api-main'
docker.api-main.stop:
	make docker.stop container='api-main'
docker.api-main.restart:
	make docker.restart container='api-main'
docker.api-main.reboot:
	make docker.reboot container='api-main'

.PHONY: docker.api-auth.build docker.api-auth.run docker.api-auth.up docker.api-auth.stop docker.api-auth.restart docker.api-auth.reboot
docker.api-auth.build:
	make docker.build container='api-auth'
docker.api-auth.run:
	make docker.run container='api-auth'
docker.api-auth.up:
	make docker.up container='api-auth'
docker.api-auth.stop:
	make docker.stop container='api-auth'
docker.api-auth.restart:
	make docker.restart container='api-auth'
docker.api-auth.reboot:
	make docker.reboot container='api-auth'

.PHONY: docker.front.build docker.front.run docker.front.up docker.front.stop docker.front.restart docker.front.reboot
docker.front.build:
	make docker.build container='front'
docker.front.run:
	make docker.run container='front'
docker.front.up:
	make docker.up container='front'
docker.front.stop:
	make docker.stop container='front'
docker.front.restart:
	make docker.restart container='front'
docker.front.reboot:
	make docker.reboot container='front'

