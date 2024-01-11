FROM golang:latest

RUN mkdir /files
COPY concurentTCP.go /files
WORKDIR /files

RUN go build -o /files/concurentTCP concurentTCP.go
ENTRYPOINT [ "/files/concurentTCP", "80" ]