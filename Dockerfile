FROM golang:alpine
RUN apk add --update make
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN make build
CMD ["./bin/goexample"]