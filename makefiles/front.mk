.PHONY: pnpm pnpm.install pnpm.run pnpm.lint
pnpm:
	make docker.front.run args="pnpm ${args}"
pnpm.install:
	make pnpm args='install'
pnpm.run:
	make pnpm args="run ${args}"
pnpm.lint:
	make pnpm.run args='lint'
