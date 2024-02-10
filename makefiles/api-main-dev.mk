.PHONY: api-main-dev api-main-dev.up api-main-dev.stop api-main-dev.restart
api-main-dev:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml ${cmd} ${opts} api-main-dev ${args}
api-main-dev.up:
	make api-main-dev cmd='up' opts='-d'
api-main-dev.stop:
	make api-main-dev cmd='stop'
api-main-dev.restart:
	make api-main-dev cmd='restart'
api-main-dev.reboot:
	make api-main-dev.stop && make api-main-dev.up
