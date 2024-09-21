WORK_DIR = $(shell pwd)

PROJECT := Infozio backend

accounts:
	go run src/accounts/cmd/main.go

attendance:
	go run src/attendance/cmd/main.go

exam:
	go run src/exam/cmd/main.go

learning:
	go run src/learning/cmd/main.go

running:
	CompileDaemon -build="go build -o ./cmd/main ./cmd" -command=./cmd/main