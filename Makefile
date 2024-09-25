WORK_DIR = $(shell pwd)

PROJECT := Infozio backend

running:
	CompileDaemon -build="go build -o ./cmd/main ./cmd" -command=./cmd/main

d-compose:
	docker-compose pull infuzio; docker-compose up

d-build:
	docker build -t abdulrahimom/infuzio .; docker push abdulrahimom/infuzio:latest

d-both:
	docker-compose pull infuzio; docker build -t abdulrahimom/infuzio .; docker push abdulrahimom/infuzio:latest; docker-compose up
