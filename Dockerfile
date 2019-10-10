FROM golang:1.13.1-alpine3.10 

WORKDIR /workspace
ENV GO111MODULE on
COPY . /workspace
RUN go build -o main main.go
CMD ["./main"]