build: 
	@go build -o bin/app.exe main.go

run: build
	@./bin/app.exe --port 5000

test: 	
	@go test -v ./...

seed:
	@go run scripts/seed.go