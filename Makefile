build: 
	@templ generate
	echo "111"
	@go build -o bin/app.exe -v main.go

run: build
	echo "111"
	@./bin/app.exe --port 5000

test: 	
	@go test -v ./... -count=1

seed:
	@go run scripts/seed.go

gop:
	@go run play/play.go
	
docker:
	docker build -t command-constructor .
	docker run --user root -p 5000:5000 command-constructor