WORK_DIR = $(shell pwd)

PROJECT := Infozio backend

default: run

run:
	go run main.go

user:
	SERVER_TYPE=user PORT=5000 go run main.go

admin:
	SERVER_TYPE=admin PORT=5000 go run main.go

quiz:
	SERVER_TYPE=quiz PORT=5000 go run main.go