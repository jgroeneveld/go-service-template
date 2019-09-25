# go-cloud-run-template

Simple template prepared to be used for services running via google cloud run using go and minimal docker image.

Final image size is < 8mb.

## Usage

- Copy repo contents
- Adjust the module in `go.mod`

## Build and run with docker

      make docker-build docker-run
      http :9000/
      
## Publish to google cloud run

TODO

## Development

Run tests

      make test

Use `https://github.com/codegangsta/gin` to watch changes

      go get github.com/codegangsta/gin
      make start

## Decisions

### Server

- `go-chi` for routing

### Testing

- `trial` for assertions
- `schema` for json matching