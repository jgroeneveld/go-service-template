# go-cloud-run-template

Simple template prepared to be used for services running via google cloud run using go and minimal docker image.

## Usage

- Adjust the module in `go.mod`
- Build with docker

      docker build -t go-cloud-run-template .
      docker run -p 9000:8080 go-cloud-run-template
      http :9000/

## Decisions

- `go-chi` for routing