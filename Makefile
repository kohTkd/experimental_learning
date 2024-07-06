args=
cmd=
container=
opts=

.SILENT:

all:
	@echo 'this is dummy'

.PHONY: setup setup.docker.mk
setup: setup.docker.mk docker.build
setup.docker.mk:
	@ruby ./tools/setup_docker_commands.rb

.PHONY: aws.exec
aws.exec:
	AWS_PROFILE=private aws ${cmd} ${args}

.PHONY: docker.cmd docker.build docker.run docker.up docker.stop docker.restart docker.reboot docker.logs docker.attach
docker.cmd:
	docker compose ${cmd} ${opts} ${container} ${args}
docker.build:
	make docker.cmd cmd='build'
docker.run:
	make docker.cmd cmd='run' opts='--rm -it'
docker.exec:
	make docker.cmd cmd='exec'
docker.up:
	make docker.cmd cmd='up' opts='-d'
docker.stop:
	make docker.cmd cmd='stop'
docker.restart:
	make docker.cmd cmd='restart'
docker.reboot:
	make docker.stop && make docker.up
docker.logs:
	make docker.cmd cmd='logs'
docker.attach:
	make docker.cmd cmd='exec' opts='-it' args='sh'
docker.down:
	make docker.cmd cmd='down'

include makefiles/*
