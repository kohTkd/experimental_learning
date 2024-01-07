.PHONY: api-dev api-dev.up api-dev.stop api-dev.restart
api-dev:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml ${cmd} ${opts} api-dev ${args}
api-dev.up:
	make api-dev cmd='up' args='-d'
api-dev.stop:
	make api-dev cmd='stop'
api-dev.restart:
	make api-dev cmd='restart'
api-dev.reboot:
	make api-dev.stop && make api-dev.up
