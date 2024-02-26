.PHONY: go go.install go.get go.clean go.run go.generate
go:
	make docker.api-auth.run args="go ${args}"
go.get:
	make go args="get -d ${args}"
go.clean:
	make go args="clean -i ${args}"
go.install:
	make go args='mod download'
go.run:
	make go args="run -mod=mod ${args}"
go.generate:
	make go args="generate ${args}"

.PHONY: go.ent go.ent.new go.ent.generate
go.ent:
	make go.run args="entgo.io/ent/cmd/ent ${args}"
go.ent.new:
	make go.ent args="new ${args}"
go.ent.generate:
	make go.generate args='./ent'

.PHONY: go.lint
go.lint:
	make docker.api-auth.run args='staticcheck ./...'
