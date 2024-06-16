.PHONY: go go.install go.get go.clean go.run go.generate
go:
	make docker.api-auth.run args="go ${args}"
go.get:
	make go args="get -d ${args}"
go.clean:
	make go args="clean -i ${args}"
go.install:
	make go args='mod download'
go.tidy:
	make go args='mod tidy'
go.run:
	make go args="run -mod=mod ${args}"
go.generate:
	make go args="generate ${args}"

.PHONY: go.ent go.ent.new go.ent.generate go.ent.migrate
go.ent:
	make go.run args="entgo.io/ent/cmd/ent ${args}"
go.ent.new:
	make go.ent args="new ${args}"
go.ent.init:
	make go.generate args='./ent'
go.ent.generate:
	make go.run args="cmd/migration/main.go ${args}"

.PHONY: go.atlas.inspect go.atlas.apply go.atlas.validate
go.atlas.inspect:
	docker compose exec api-auth sh -c 'atlas schema inspect -u "mysql://$${DB_WRITE_USER}:$${DB_WRITE_PASSWORD}@$${DB_WRITE_HOST}/$${DB_NAME}" > atlas/schema.hcl'
go.atlas.apply:
	docker compose exec api-auth sh -c 'atlas migrate apply --dir "file://atlas/migrations" -u "mysql://$${DB_WRITE_USER}:$${DB_WRITE_PASSWORD}@$${DB_WRITE_HOST}/$${DB_NAME}"'
	make go.atlas.inspect
go.atlas.validate:
	docker compose exec api-auth sh -c 'atlas migrate validate --dir file://atlas/migration'


.PHONY: go.lint
go.lint:
	make docker.api-auth.run args='staticcheck ./...'
