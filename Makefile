build:
	@rm -rf .commitix;\
	go build -o bin/commitix cmd/commitix/main.go;