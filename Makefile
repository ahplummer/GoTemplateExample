.PHONY: default
default: displayhelp ;

displayhelp:
	@echo Use "clean, showcoverage, tests, build, buildlinux or run" with make, por favor.

showcoverage: tests
	@echo Running Coverage output
	go tool cover -html=coverage.out

tests: clean
	@echo Running Tests
	go test --coverprofile=coverage.out ./...

docker:
	docker build -t goexample:latest . -f Dockerfile
	docker run -it --env PORT=3333 --env RUN_LOCAL=TRUE -p 80:3333 goexample:latest

run: build
	@echo Running program
	RUN_LOCAL=true ./bin/goexample

build: clean
	@echo Running build command
	go build -o bin/goexample src/main.go src/website.go

buildlinux:
	env GOOS=linux go build -ldflags="-s -w" -o bin/goexample src/main.go src/website.go

clean:
	@echo Removing binary TODO
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
