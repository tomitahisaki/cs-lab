PKGS := $(shell go list ./... | grep -v /vendor/)
CHANGED := $(shell git diff --name-only origin/$(BASE_REF)...HEAD 2>/dev/null | cut -d/ -f1 | sort -u | grep -E '^(cs|lab|systems|algo|shared)$$' || true)
BASE_REF ?= main

.PHONY: test lint bench tidy test-changed
test:
	go test -race ./...

lint:
	golangci-lint run ./...

bench:
	go test -bench=. -benchmem ./...

tidy:
	for m in cs lab systems algo shared; do (cd $$m && go mod tidy); done

test-changed:
	@if [ -z "$(CHANGED)" ]; then echo "no changes in modules"; exit 0; fi; \
	for m in $(CHANGED); do (cd $$m && go test -race ./...); done
