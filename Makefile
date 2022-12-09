BIN=main
build:
	go build -o $(BIN) main.go
run: build
	./$(BIN)
