FROM golang:1.19-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go *.html ./
RUN go build -o server

EXPOSE 8080
CMD ["./server"]