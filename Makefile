TEST?=./...

default: test

build:
	@sh -c "'$(CURDIR)/scripts/build.sh' app"

lint:
	@sh -c "'$(CURDIR)/scripts/lint.sh'"

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/fmt.sh'"

covercheck:
	@sh -c "'$(CURDIR)/scripts/coverage.sh' 90"
	rm coverage.out

cover:
	@sh -c "'$(CURDIR)/scripts/cover.sh'"

test: fmtcheck
	@sh -c "'$(CURDIR)/scripts/test.sh'"

