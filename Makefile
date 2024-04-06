.PHONY: dev build

dev:
	@templ generate --watch

build:
	@templ generate

.DEFAULT_GOAL := dev