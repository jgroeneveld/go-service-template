PROJECTNAME := $(shell basename "$(shell pwd)")

test:
	go test ./...

start:
	~/go/bin/gin --bin bin/gin-bin run main.go

docker-build:
	docker build -t $(PROJECTNAME) .

docker-run:
	docker run -p 9000:8080 $(PROJECTNAME)