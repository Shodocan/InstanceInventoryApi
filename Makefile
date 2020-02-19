default: test

build: covercheck
	@sh -c "'$(CURDIR)/scripts/build.sh' bin/app"

covercheck: cover
	@sh -c "'$(CURDIR)/scripts/coverage.sh' 90"

cover: test
	@sh -c "'$(CURDIR)/scripts/cover.sh'"

test: lint
	@sh -c "'$(CURDIR)/scripts/test.sh'"

lint:
	@sh -c "'$(CURDIR)/scripts/lint.sh'"

release:
	@sh -c "'$(CURDIR)/scripts/release.sh' local"

run:
	@sh -c "'$(CURDIR)/scripts/run.sh'"