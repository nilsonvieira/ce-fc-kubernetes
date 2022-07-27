FROM golang:latest
RUN mkdir -p src/goapp
WORKDIR src/goapp
COPY . .
RUN go mod init
RUN go build -o server
CMD ["./server"]