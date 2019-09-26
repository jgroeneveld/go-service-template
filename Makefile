PROJECTNAME := $(shell basename "$(shell pwd)")
GIN := ~/go/bin/gin

test:
	@cd src && go test ./...

start:
	@cd src && $(GIN) --bin ../bin/gin-bin run main.go

clean:
	rm ./bin/*

docker-build:
	docker build -t $(PROJECTNAME) .

docker-run:
	docker run -p 9000:8080 $(PROJECTNAME)