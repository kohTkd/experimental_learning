.PHONY: docker.api.build docker.api.run docker.api.up docker.api.stop docker.api.restart docker.api.reboot
docker.api.build:
	make docker.build container='api'
docker.api.run:
	make docker.run container='api'
docker.api.up:
	make docker.up container='api'
docker.api.stop:
	make docker.stop container='api'
docker.api.restart:
	make docker.restart container='api'
docker.api.reboot:
	make docker.reboot container='api'

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

