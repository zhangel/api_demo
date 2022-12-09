BIN=import_db
build:
	go build -o $(BIN) main.go
run: build
	./$(BIN) convert --src=./ --dest=./ --file_type=json,xml
