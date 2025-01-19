build: 
	@templ generate
	@go build -o bin/app.exe -v main.go

run: build
	@./bin/app.exe --port 5000

test: 	
	@go test -v ./... -count=1

seed:
	@go run scripts/seed.go

gop:
	@go run play/play.go
	
docker:
	docker build -t command-constructor .
	docker run -p 5000:5000 command-constructor