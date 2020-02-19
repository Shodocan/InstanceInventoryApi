TEST?=./...

default: test

build:
	@sh -c "'$(CURDIR)/scripts/build.sh' app"

lint:
	@sh -c "'$(CURDIR)/scripts/lint.sh'"

covercheck: cover
	@sh -c "'$(CURDIR)/scripts/coverage.sh' 90"

test: lint
	@sh -c "'$(CURDIR)/scripts/test.sh'"

cover: test
	@sh -c "'$(CURDIR)/scripts/cover.sh'"