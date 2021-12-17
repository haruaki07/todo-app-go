# run backend service
run:
	@make build
	@./bin/todoapp.exe

build:
	@go build -v -o bin/todoapp.exe backend/main.go
