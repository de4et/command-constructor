build: 
	@go build -o bin/app.exe main.go

run: build
	templ generate
	@./bin/app.exe --port 5000

test: 	
	@go test -v ./... -count=1

seed:
	@go run scripts/seed.go

gop:
	@go run play/play.go
	