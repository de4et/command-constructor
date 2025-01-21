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
	docker tag command-constructor:latest de4et/command-constructor:latest

docker-push: 
	docker push de4et/command-constructor:latest

docker-pull:
	docker pull de4et/command-constructor:latest

docker-start:
	docker compose up --detach

docker-stop:
	docker compose down

send:
	pscp -r ./ root@94.159.104.84:/root/
