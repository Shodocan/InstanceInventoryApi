TEST?=./...

default: test

build: covercheck
	@sh -c "'$(CURDIR)/scripts/build.sh' app"

covercheck: cover
	@sh -c "'$(CURDIR)/scripts/coverage.sh' 90"

cover: test
	@sh -c "'$(CURDIR)/scripts/cover.sh'"

test: lint
	@sh -c "'$(CURDIR)/scripts/test.sh'"

lint:
	@sh -c "'$(CURDIR)/scripts/lint.sh'"
