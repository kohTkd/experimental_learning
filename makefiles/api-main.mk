.PHONY: bundle bundle.install bundle.exec
bundle:
	make docker.api-main.run args="bundle ${args}"
bundle.install:
	make bundle args='install'
bundle.exec:
	make bundle args="exec ${args}"

.PHONY: rubocop rubocop.fix rspec
rubocop:
	make bundle.exec args='rubocop'
rubocop.fix:
	make bundle.exec args='rubocop -A'
rspec:
	make bundle.exec cmd='rspec'

.PHONY: rails rails.console rails.routes rails.pry
rails:
	make bundle.exec args="rails ${args}"
rails.console:
	make rails args='c'
rails.routes:
	make rails args='routes'
rails.pry:
	docker attach `docker-compose ps -q api`

## Rails生成関連
.PHONY: rails.generate rails.generate.model rails.generate.controller rails.generate.migration \
				rails.generate.worker rails.generate.rspec.request \
				rails.pry
rails.generate:
	make rails args="g ${args}"
rails.generate.model:
	make rails.generate args='model'
rails.generate.controller:
	make rails.generate args='controller'
rails.generate.migration:
	make rails.generate args='migration'
rails.generate.worker:
	make rails.generate args='sidekiq:worker'
rails.generate.mailer:
	make rails.generate args='mailer'
rails.generate.rspec.request:
	make rails.generate args='rspec:request'

# Rails DB関連
.PHONY: rails.db.create rails.db.migrate rails.db.seed rails.db.prepare
rails.db.create:
	make bundle.exec cmd='rails db:create > /dev/null'
rails.db.migrate:
	make bundle.exec cmd='rails db:migrate'
rails.db.rollback:
	make bundle.exec cmd='rails db:rollback'
rails.db.seed:
	make bundle.exec cmd='rails db:seed > /dev/null'
rails.db.prepare:
	make rails.db.create
	make rails.db.migrate
	make rails.db.seed
