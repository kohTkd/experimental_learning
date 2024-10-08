.PHONY: docker.api-main.build docker.api-main.run docker.api-main.up docker.api-main.stop docker.api-main.restart docker.api-main.reboot docker.api-main.logs docker.api-main.attach docker.api-main.down
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
docker.api-main.logs:
	make docker.logs container='api-main'
docker.api-main.attach:
	make docker.attach container='api-main'
docker.api-main.down:
	make docker.down container='api-main'

.PHONY: docker.api-auth.build docker.api-auth.run docker.api-auth.up docker.api-auth.stop docker.api-auth.restart docker.api-auth.reboot docker.api-auth.logs docker.api-auth.attach docker.api-auth.down
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
docker.api-auth.logs:
	make docker.logs container='api-auth'
docker.api-auth.attach:
	make docker.attach container='api-auth'
docker.api-auth.down:
	make docker.down container='api-auth'

.PHONY: docker.front.build docker.front.run docker.front.up docker.front.stop docker.front.restart docker.front.reboot docker.front.logs docker.front.attach docker.front.down
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
docker.front.logs:
	make docker.logs container='front'
docker.front.attach:
	make docker.attach container='front'
docker.front.down:
	make docker.down container='front'

.PHONY: docker.database-main.build docker.database-main.up docker.database-main.stop docker.database-main.logs docker.database-main.down docker.database-main.attach
docker.database-main.build:
	make docker.build container='database-main'
docker.database-main.up:
	make docker.up container='database-main'
docker.database-main.stop:
	make docker.stop container='database-main'
docker.database-main.logs:
	make docker.logs container='database-main'
docker.database-main.down:
	make docker.down container='database-main'
docker.database-main.attach:
	make docker.attach container='database-main'

.PHONY: docker.database-auth.build docker.database-auth.up docker.database-auth.stop docker.database-auth.logs docker.database-auth.down docker.database-auth.attach
docker.database-auth.build:
	make docker.build container='database-auth'
docker.database-auth.up:
	make docker.up container='database-auth'
docker.database-auth.stop:
	make docker.stop container='database-auth'
docker.database-auth.logs:
	make docker.logs container='database-auth'
docker.database-auth.down:
	make docker.down container='database-auth'
docker.database-auth.attach:
	make docker.attach container='database-auth'
