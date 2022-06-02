dir=`pwd`
.PHONY: build
build:
	@docker build -f docker/Dockerfile --build-arg "PWD=$(dir)" -t ria-course-crud .
.PHONY: run
run:
	@docker-compose -f ./docker/compose/dev.yaml -f ./docker/compose/local.yaml  up

