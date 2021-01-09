# This is a bootstrappable Golang Project that:
* Is Dockerized
* Has Tests
* Can be deployed to AWS Lambda
* Can be run locally
* Has a swanky Makefile
* Uses HTML Templates in Go

## Instructions
* Runnable as a Lambda (in AWS)...(You'll want to `aws configure` first.)
```
make buildlinux
sls deploy
```
* Runnable Locally
```
make run
```
* Runnable in docker
```
make docker
```
* Run Tests
```
make showcoverage
```

